<template>
  <div class="terminal" id="terminal-container" ref="terminalRef"></div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import 'xterm/css/xterm.css'
import axios from 'axios'

const route = useRoute()
const host = route.query.host
const cid = route.query.cid
const command = route.query.command
const port = route.query.port

const ws = ref<WebSocket>()
const term = ref<Terminal>()
const fitAddon = new FitAddon()

const terminalRef = ref()
const base = `http://${host}:${port}/v1/terminal/resize?cid=${cid}&command=${command}`

const resizeRemoteTerminal = () => {
  fitAddon.fit()
  const { cols, rows } = term.value as Terminal
  console.log(cols, rows)
  var url = base + `&w=${cols}&h=${rows}`
  axios.post(url)
}

const createWs = () => {
  var url = base + `&w=200&h=30`
  axios.post(url)
  ws.value = new WebSocket(
    `ws://${host}:${port}/v1/terminal?container=${cid}&command=${command}`
  )
  ws.value.onmessage = (e: any) => {
    term.value!.write(e.data)
  }
}

const initTerm = () => {
  term.value = new Terminal({
    rendererType: 'canvas',
    lineHeight: 1.2,
    fontSize: 16,
    cursorBlink: true,
    cursorStyle: 'underline'
  })
  term.value.open(terminalRef.value)
  term.value.focus()
  term.value.loadAddon(fitAddon)
  term.value.onData((data) => ws.value?.send(data))
  setTimeout(() => {
    fitAddon.fit()
  }, 100)
}

const onTerminalResize = () => {
  window.addEventListener('resize', resizeRemoteTerminal)
}

onMounted(() => {
  createWs()
  initTerm()
  onTerminalResize()
})
</script>

<style lang="less">
.terminal {
  width: 100%;
  height: 100%;

  .xterm-viewport {
    width: 100%;
    height: 100%;
  }
}
</style>
