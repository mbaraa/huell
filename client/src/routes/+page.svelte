<script lang="ts">
    let downloaded = 0;
    let chunkSize = 1000 * 1000 * 100;
    const kiloByte = 1000;
    const megaByte = kiloByte * 1000;
    const gigaByte = megaByte * 1000;
    const terraByte = gigaByte * 1000;

    function formatSize(size: number): string {
        let sizeUnit = ( (size >= terraByte)? "TBs": (size >= gigaByte)? "GBs": (size >= megaByte)? "MBs": (size >= kiloByte)? "KBs": "Bytes" );

        switch (sizeUnit) {
        case "Bytes":
            break;
        case "KBs":
        size /= 1000;
            break;
        case "MBs":
        size /= 1000*1000;
            break;
        case "GBs":
        size /= 1000*1000*1000;
            break;
        case "TBs":
        size /= 1000*1000*1000*1000;
            break;
        }

        return `${size} ${sizeUnit}`
    }

    let interval = null;

     async function startZeOperation() {
        interval = setInterval( async () => {
            await fetch(`http://localhost:9090/meow?fs=${chunkSize}`)
            .then(resp => resp.text())
            .then(text => {
                downloaded += chunkSize;
            })
            }, 128);
    }

    async function stopZeOperation() {
        clearInterval(interval);
    }

</script>

<div>
    <label for="size">{`Enter chunk size (in bytes) default is ${formatSize(chunkSize)}`}</label>
    <input id="size" type="number" bind:value={chunkSize} />
    <br/>
    <span>Downloaded {formatSize(downloaded)} </span>
    <br/>
    <button on:click={startZeOperation}>Start</button>
    <button on:click={stopZeOperation}>Stop</button>

</div>
