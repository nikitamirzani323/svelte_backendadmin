<script>
    import { createEventDispatcher } from "svelte";
    import Panel_Table from "../components/panel_table.svelte"
    export let panel_page =  "";
    export let panel_total =  0;
    export let panel_height = "550px";
    export let panel_button_new =  false;
    export let panel_button_refresh =  false;

    let dispatch = createEventDispatcher();

    const HandleNewForm = () => {
        dispatch("handleNewForm", "call");
    };
    const HandleRefresh = () => {
        dispatch("handleRefresh", "call");
    };
</script>
<div class="container mx-auto px-2 lg:px-28 ">
    <div class="bg-white shadow-lg p-5 ">
        <div class="flex flex-col gap-2">
            <div class="flex">
                <h1 class="text-slate-600 font-bold text-sm lg:text-3xl uppercase w-full">{panel_page}</h1>
                <div class="hidden sm:flex md:flex justify-end w-full gap-2 ">
                    {#if panel_button_new}
                    <button 
                        on:click={() => {
                            HandleNewForm();
                        }}
                        class="btn bg-primary hover:bg-primary  rounded-md shadow-lg shadow-[#6eb5d8] border-none  ">New</button>
                    {/if}
                    {#if panel_button_refresh}
                    <button on:click={() => {
                            HandleRefresh();
                        }} class="btn btn-secondary hover:bg-secondary shadow-lg shadow-[#c76797]  rounded-md lg:inline-flex">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                        </svg>
                    </button>
                    {/if}
                </div>
            </div>
            <div class="relative w-full">
                <slot name="panel_search" />
            </div>
            <Panel_Table panel_height="{panel_height}">
                <slot:template slot="paneltable_body">
                    <slot name="panel_body" />
                </slot:template>
            </Panel_Table>

            <div class="bg-[#F7F7F7] rounded-sm h-16 p-5">
                <span class="font-semibold">TOTAL ROW : {panel_total}</span>
            </div>
        </div>
    </div>
</div>