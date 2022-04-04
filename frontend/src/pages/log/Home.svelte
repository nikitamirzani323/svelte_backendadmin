<script>
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Input_custom from '../../components/Input.svelte' 
    import Modal_alert from '../../components/Modal_alert.svelte' 

    export let path_api = "";
    export let token = "";
    export let listHome = [];
    export let totalrecord = 0;

    let page = "Log Management";
    let isModalLoading = false
    let isModalNotif = false
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let msg_error = "";
    
    let searchHome = "";
    let filterHome = [];
    let dispatch = createEventDispatcher();
    
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    
    
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_username
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) ||
                    item.home_page
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) 
            );
        } else {
            filterHome = [...listHome];
        }
    }
</script>
<div class="{loader_class} w-40 fixed top-1 left-[45%] right-0 z-50">
    <div class="alert alert-warning shadow-lg rounded-md bg-[#FFF6BF]">
        <div>
          <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current flex-shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" /></svg>
          <span>{loader_msg}</span>
        </div>
    </div>
</div>
<div class="container mx-auto bg-white shadow-lg p-5">
    <div class="flex flex-col gap-2">
        <div class="flex">
            <h1 class="text-black font-bold text-sm lg:text-4xl uppercase w-full">{page}</h1>
            <div class="hidden sm:flex md:flex justify-end w-full gap-2 ">
                <button on:click={() => {
                    RefreshHalaman();
                    }} class="btn btn-primary rounded-md lg:inline-flex">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                    </svg>
                </button>
            </div>
        </div>
        <div class="relative w-full">
            <div class="absolute inset-y-0 left-0 flex items-center pl-2">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 stroke-current text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
            </div>
            <input 
                bind:value={searchHome}
                type="text" placeholder="Search by Username, Page" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 ">
        </div>
        
        <div class="hidden sm:inline w-full scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[550px] overflow-y-scroll">
            <table class="table table-compact w-full">
                <thead class="sticky top-0">
                    <tr>
                        <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">NO</th>
                        <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">DATETIME</th>
                        <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">USERNAME</th>
                        <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">PAGE</th>
                        <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">ACTION</th>
                        <th width="*" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">NOTE</th>
                    </tr>
                </thead>
                {#if filterHome != ""}
                    <tbody>
                        {#each filterHome as rec}
                        <tr>
                            <td class="text-xs lg:text-sm align-top text-center">{rec.home_no}</td>
                            <td class="text-xs lg:text-sm align-top text-center">{rec.home_datetime}</td>
                            <td class="text-xs lg:text-sm align-top text-left">{rec.home_username}</td>
                            <td class="text-xs lg:text-sm align-top text-left">{rec.home_page}</td>
                            <td class="text-xs lg:text-sm align-top text-left">{rec.home_tipe}</td>
                            <td class="text-xs lg:text-sm align-top text-left">{@html rec.home_note}</td>
                        </tr>
                        {/each}
                    </tbody>
                {:else}
                    <tbody>
                        <tr>
                            <td colspan="6" class="text-center">
                                <progress class="self-start progress progress-primary w-56"></progress>
                            </td>
                        </tr>
                    </tbody>
                {/if}
            </table>
            
        </div>

        <div class="bg-[#F7F7F7] rounded-sm h-16 p-5">
        <span class="font-semibold">TOTAL ROW : {totalrecord}</span>
        </div>
    </div>
</div>





<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />



