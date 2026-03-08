<script lang="ts">
    import { onMount } from 'svelte';

    const width = 600;
    const height = 600;

    let canvas: any;
    let context: any;
    let isDrawing: boolean;
    let start: { x: number; y: number };

    let t: number, l: number;

    onMount(() => {
        context = canvas.getContext('2d');
        context.lineWidth = 3;

        context.strokeStyle = '#fff';

        handleSize();
    });

    const handleStart = ({ offsetX: x, offsetY: y }: { offsetX: number; offsetY: number }) => {
        isDrawing = true;
        start = { x, y };
    };

    const handleEnd = () => {
        isDrawing = false;
    };

    const handleMove = ({ offsetX: x1, offsetY: y1 }: { offsetX: number; offsetY: number }) => {
        if (!isDrawing) return;

        const { x, y } = start;
        context.beginPath();
        context.moveTo(x, y);
        context.lineTo(x1, y1);
        context.closePath();
        context.stroke();

        start = { x: x1, y: y1 };
    };

    const handleSize = () => {
        const { top, left } = canvas.getBoundingClientRect();
        t = top;
        l = left;
    };
</script>

<svelte:window on:resize={handleSize} />

<canvas
    {width}
    {height}
    bind:this={canvas}
    onmousedown={handleStart}
    ontouchstart={(e) => {
        const { clientX, clientY } = e.touches[0];
        handleStart({
            offsetX: clientX - l,
            offsetY: clientY - t
        });
    }}
    onmouseup={handleEnd}
    ontouchend={handleEnd}
    onmouseleave={handleEnd}
    onmousemove={handleMove}
    ontouchmove={(e) => {
        const { clientX, clientY } = e.touches[0];
        handleMove({
            offsetX: clientX - l,
            offsetY: clientY - t
        });
    }}
></canvas>

<style>
    canvas {
        background: var(--bg-color);
        border: dashed 2px var(--main-border-color);
    }
</style>
