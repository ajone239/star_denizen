<script lang="ts">
    import Canvas from '$lib/components/Canvas.svelte';
    import { onMount } from 'svelte';

    let socket: WebSocket | null = null;

    let message = $state('');
    let messages: string[] = $state([]);

    function sendMessage() {
        if (socket === null) {
            alert('attempted to send with no websocket');
            return;
        }

        socket.send(message);

        messages.push('Sent: ' + message);

        message = '';
    }

    onMount(() => {
        const protocol = window.location.protocol === 'https:' ? 'wss://' : 'ws://';
        const hostname = window.location.host;
        const wsPath = '/api/ws';

        let wsUrl = `${protocol}${hostname}${wsPath}`;

        console.log(wsUrl);

        socket = new WebSocket(wsUrl);

        socket.onopen = function () {
            messages.push('Opened!');
        };

        socket.onmessage = function (event) {
            messages.push('Recv: ' + event.data);
        };

        socket.onclose = function () {
            messages.push('Closed!');
        };
    });
</script>

<h2>Echo</h2>
<input type="text" bind:value={message} placeholder="Enter message" />
<button onclick={sendMessage}>Send</button>
<pre>
{#each messages as msg}
        {msg}
    {/each}
</pre>

<Canvas />
