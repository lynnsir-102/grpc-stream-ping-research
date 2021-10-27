# GRPC "too many pings from client" 排查

## 背景

Rainbow3.0 于近日阶段性替换上线, 适逢大促运维侧进行批量机器扩容, 发现集中扩容时,
偶现 too many pings from client, 同时断开 stream, 由 Rainbow 侧重连机制重新连接.

## 排查

```go
func (t *http2Server) handlePing(f *http2.PingFrame) {
 if f.IsAck() {
  if f.Data == goAwayPing.data && t.drainChan != nil {
   close(t.drainChan)
   return
  }
  // Maybe it's a BDP ping.
  if t.bdpEst != nil {
   t.bdpEst.calculate(f.Data)
  }
  return
 }
 pingAck := &ping{ack: true}
 copy(pingAck.data[:], f.Data[:])
 t.controlBuf.put(pingAck)

 now := time.Now()
 defer func() {
  t.lastPingAt = now
 }()
 // A reset ping strikes means that we don't need to check for policy
 // violation for this ping and the pingStrikes counter should be set
 // to 0.
 if atomic.CompareAndSwapUint32(&t.resetPingStrikes, 1, 0) {
  t.pingStrikes = 0
  return
 }
 t.mu.Lock()
 ns := len(t.activeStreams)
 t.mu.Unlock()
 if ns < 1 && !t.kep.PermitWithoutStream {
  // Keepalive shouldn't be active thus, this new ping should
  // have come after at least defaultPingTimeout.
  if t.lastPingAt.Add(defaultPingTimeout).After(now) {
   t.pingStrikes++
  }
 } else {
  // Check if keepalive policy is respected.
  if t.lastPingAt.Add(t.kep.MinTime).After(now) {
   t.pingStrikes++
  }
 }

 if t.pingStrikes > maxPingStrikes {
  // Send goaway and close the connection.
  if logger.V(logLevel) {
   logger.Errorf("transport: Got too many pings from the client, closing the connection.")
  }
  t.controlBuf.put(&goAway{code: http2.ErrCodeEnhanceYourCalm, debugData: []byte("too_many_pings"), closeConn: true})
 }
}
```

发现客户端在发起 keeplive ping 包超过服务端 EnforcementPolicy MinTime 设置时, 触发该异常

## 试验

通过 server/client 示例进行调试, </br>

| 序号 | client: ClientParameters | server: ServerParameters | server: EnforcementPolicy | 现象 |
| ---- | ---- | ---- | ---- | ---- |
| 1 | Time : 12 * time.Second | Time: 20 * time.Second | MinTime : 10 * time.Second | 每12s client wrote PING, stream持续正常 (取客户端较小ping周期) |
| 2 | Time : 12 * time.Second | Time: 20 * time.Second | MinTime : 15 * time.Second | 每12s client wrote PING (取客户端较小ping周期, server wrote GOAWAY, 断开重连 |
| 3 | Time : 20 * time.Second | Time: 15 * time.Second | MinTime : 10 * time.Second | 每15s server wrote PING (取服务端较小ping周期), stream持续正常 |
| 4 | Time : 10 * time.Second | Time: 10 * time.Second | MinTime : 10 * time.Second | 每10s server/clent 交替wrote PING, stream持续正常 |

## 结论

结合试验, 目前 client/server keeplive 配置均为10s, server 端 enforcementpolicy 也为10s, 会在空闲时交替发送 ping 包, 在服务端受到压力时, 服务端可能存在不发送 ping 包情况, 同时接收处理客户端前后2个 ping 包出现延迟, 判断小于 policy mintime, 即可能出现当前情况, 可以调整 enforcementpolicy, mintime 从 10s 调整为 7s (10-3)

## 参考资料

<https://pandaychen.github.io/2020/09/01/GRPC-CLIENT-CONN-LASTING/>
