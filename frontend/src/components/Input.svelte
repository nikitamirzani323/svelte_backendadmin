<script>
    export let value
    export let input_onchange;
    export let input_invalid = false;
    export let input_required = false;
    export let input_autofocus = false;
    export let input_enabled = true;
    export let input_tipe = "text";
    export let input_attr = "";
    export let input_id = "";
    export let input_placeholder = "";
    let input_class = "peer w-full rounded px-3  border border-gray-300  focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none placeholder-transparent";
    let input_datetime_class = "peer w-full text-center rounded px-3  border border-gray-300  focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none placeholder-transparent";
    let input_number_class = "peer w-full text-right rounded px-3  border border-gray-300  focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none placeholder-transparent";
    let show_eye = true;
    function handleShowEye()  {
        show_eye = !show_eye;
    }
    function typeAction(node){
        node.type = input_attr
    }
    const handleKeyboard_number = (e) => {
        if (isNaN(parseInt(e.target.value))) {
            return e.target.value = "";
        }else{
            return e.target.value = parseInt(e.target.value);
        }
	}
    const handleKeyboard_number_blur = (e) => {
        return e.target.value = parseInt(e.target.value);
	}
</script>
{#if input_tipe == "text"}
    <div class="relative form-control">
        {#if input_autofocus}
        <input
            autofocus
            on:change={input_onchange}
            bind:value
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
                    bind:value
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
                    bind:value
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
        <label for="username" class="absolute left-3 top-[-0.5rem]  text-xs cursor-text 
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
        {#if show_eye}
            <input
                use:typeAction 
                on:change={input_onchange}
                bind:value
                invalid={input_invalid}
                id="{input_id}"
                name="{input_id}"
                placeholder="{input_placeholder}"
                autocomplete="off"
                class="{input_class}"> 
        {:else}
            <input
                on:change={input_onchange}
                bind:value
                invalid={input_invalid}
                type="text"
                id="{input_id}"
                name="{input_id}"
                placeholder="{input_placeholder}"
                autocomplete="off"
                class="{input_class}"> 
        {/if}
        <label for="username" class="absolute left-3 top-[-0.5rem]  text-xs cursor-text 
                transition-all
                peer-placeholder-shown:text-base 
                peer-placeholder-shown:text-gray-400 
                peer-placeholder-shown:top-3 
                peer-focus:text-[#1a73e8]
                peer-focus:bg-[#fff]
                peer-focus:text-[.75rem]
                text-[#1a73e8] 
                bg-[#fff]">{input_placeholder}{input_required ? "*":""}</label>
        {#if show_eye}
            <svg on:click={() => {
                    handleShowEye();
                }}  xmlns="http://www.w3.org/2000/svg" class="absolute h-6 w-6 right-3 top-3 cursor-pointer " fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
            </svg>
        {:else}
            <svg on:click={() => {
                    handleShowEye();
                }}  xmlns="http://www.w3.org/2000/svg" class="absolute h-6 w-6 right-3 top-3 cursor-pointer" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" />
            </svg>
        {/if}
    </div> 
{/if}
{#if input_tipe == "time"}
    <div class="relative form-control">
        {#if input_enabled}        
            <input
                on:change={input_onchange}
                bind:value
                invalid={input_invalid}
                type="text" 
                id="{input_id}"
                name="{input_id}"
                placeholder="{input_placeholder}"
                autocomplete="off"
                class="{input_datetime_class}"> 
        {:else}
            <input
                on:change={input_onchange}
                bind:value
                invalid={input_invalid}
                type="text" 
                id="{input_id}"
                name="{input_id}"
                placeholder="{input_placeholder}"
                autocomplete="off"
                disabled
                class="{input_datetime_class}"> 
        {/if}
        <label for="username" class="absolute left-3 top-[-0.5rem]  text-xs cursor-text 
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
{#if input_tipe == "number_nolabel"}
    <div class="relative form-control">
        {#if input_enabled}        
            <input
                on:change={input_onchange}
                on:keyup={handleKeyboard_number}
                bind:value
                type="text" 
                id="{input_id}"
                name="{input_id}"
                placeholder="{input_placeholder}"
                autocomplete="off"
                class="{input_number_class}"> 
        {:else}
            <input
                bind:value
                type="text" 
                id="{input_id}"
                name="{input_id}"
                placeholder="{input_placeholder}"
                autocomplete="off"
                disabled
                class="{input_number_class}"> 
        {/if}
        <label for="{input_id}" class="absolute left-3 top-[-0.5rem]  text-xs cursor-text 
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
{#if input_tipe == "number"}
    <div class="relative form-control">
        {#if input_enabled}        
            <input
                on:keyup={handleKeyboard_number}
                on:blur={handleKeyboard_number_blur}
                bind:value
                type="text" 
                id="{input_id}"
                name="{input_id}"
                placeholder="{input_placeholder}"
                autocomplete="off"
                class="{input_number_class}"> 
        {:else}
            <input
                bind:value
                type="text" 
                id="{input_id}"
                name="{input_id}"
                placeholder="{input_placeholder}"
                autocomplete="off"
                disabled
                class="{input_number_class}"> 
        {/if}
        <label for="{input_id}" class="absolute left-3 top-[-0.5rem]  text-xs cursor-text 
                transition-all
                peer-placeholder-shown:text-base 
                peer-placeholder-shown:text-gray-400 
                peer-placeholder-shown:top-3 
                peer-focus:text-[#1a73e8]
                peer-focus:bg-[#fff]
                peer-focus:text-[.75rem]
                text-[#1a73e8] 
                bg-[#fff]">{input_placeholder}{input_required ? "*":""}</label>
        <span class="text-[10px] text-right text-amber-600">{new Intl.NumberFormat().format(value)}</span>
    </div>
{/if}