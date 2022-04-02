<script>
    export let input_value = "";
    export let input_onchange;
    export let input_invalid = false;
    export let input_required = false;
    export let input_autofocus = false;
    export let input_enabled = true;
    export let input_tipe = "text";
    export let input_id = "";
    export let input_placeholder = "";
    let input_class = "peer w-full rounded px-3  border border-gray-300  focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none placeholder-transparent";
</script>
{#if input_tipe == "text"}
    <div class="relative form-control">
        {#if input_autofocus}
        <input
            autofocus
            on:change={input_onchange}
            bind:value={input_value}
            invalid={input_invalid}
            type="text" 
            id="{input_id}"
            name="{input_id}"
            placeholder="{input_placeholder}"
            autocomplete="off"
            class="{input_class}"> 
        {:else}
            {#if input_enabled}        
                <input
                    on:change={input_onchange}
                    bind:value={input_value}
                    invalid={input_invalid}
                    type="text" 
                    id="{input_id}"
                    name="{input_id}"
                    placeholder="{input_placeholder}"
                    autocomplete="off"
                    class="{input_class}"> 
            {:else}
                <input
                    on:change={input_onchange}
                    bind:value={input_value}
                    invalid={input_invalid}
                    type="text" 
                    id="{input_id}"
                    name="{input_id}"
                    placeholder="{input_placeholder}"
                    autocomplete="off"
                    disabled
                    class="{input_class}"> 
            {/if}
        {/if}
        <label for="username" class="absolute left-3 top-[-0.7rem]  text-sm cursor-text 
                transition-all
                peer-placeholder-shown:text-base 
                peer-placeholder-shown:text-gray-400 
                peer-placeholder-shown:top-3 
                peer-focus:text-[#1a73e8]
                peer-focus:bg-[#fff]
                peer-focus:text-[.75rem]
                text-[#1a73e8] 
                bg-[#fff]">{input_placeholder}{input_required ? "*":""}</label>
    </div>
{/if}
{#if input_tipe == "password"}
    <div class="relative form-control">
        <input
            on:change={input_onchange}
            bind:value={input_value}
            invalid={input_invalid}
            type="password" 
            id="{input_id}"
            name="{input_id}"
            placeholder="{input_placeholder}"
            autocomplete="off"
            class="{input_class}"> 
        <label for="username" class="absolute left-3 top-[-0.7rem]  text-sm cursor-text 
                transition-all
                peer-placeholder-shown:text-base 
                peer-placeholder-shown:text-gray-400 
                peer-placeholder-shown:top-3 
                peer-focus:text-[#1a73e8]
                peer-focus:bg-[#fff]
                peer-focus:text-[.75rem]
                text-[#1a73e8] 
                bg-[#fff]">{input_placeholder}{input_required ? "*":""}</label>
    </div>
{/if}