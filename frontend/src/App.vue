<script setup>
import { RouterLink, RouterView } from "vue-router";
import HelloWorld from "./components/HelloWorld.vue";
import "xterm/css/xterm.css"; // DO NOT forget importing xterm.css
import { Terminal } from 'xterm';
import { AttachAddon } from 'xterm-addon-attach'

const term = new Terminal({
  theme: {
    background: "#202B33",
    foreground: "#F5F8FA"
  },
  rows: 30,
  cols: 80,
  screenKeys: true,
  useStyle: true,
  cursorBlink: true,
  fullscreenWin: true,
  maximizeWin: true,
  screenReaderMode: true,
},
);
window.addEventListener('load', function () {
  term.open(document.getElementById('terminal'))
  // term.write("shell> ")
  var protocol = (location.protocol === "https:") ? "wss://" : "ws://";
  var url = protocol + location.hostname + ":8376" + "/terminal"
  var ws = new WebSocket(url);
  // This addon is to connect to our websocket and send data there
  var attachAddon = new AttachAddon(ws);
  ws.onopen = function() {
    term.loadAddon(attachAddon);
    term._initialized = true;
    term.focus();
    setTimeout(function() {fitAddon.fit()});
    term.onResize(function(event) {
      var rows = event.rows;
      var cols = event.cols;
      var size = JSON.stringify({cols: cols, rows: rows + 1});
      var send = new TextEncoder().encode("\x01" + size);
      console.log('resizing to', size);
      ws.send(send);
    });
    term.onTitleChange(function(event) {
      console.log(event);
    });
    window.onresize = function() {
      fitAddon.fit();
    };
  };
})
</script>

<template>
  <!-- <header>
    <img
      alt="Vue logo"
      class="logo"
      src="@/assets/logo.svg"
      width="125"
      height="125"
    />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />
      
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header> -->

  <!-- <RouterView /> -->
  <div id = "terminal"></div>
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
