<script>
    import { createEventDispatcher } from "svelte";
    import Loader from '../../components/Loader.svelte' 
    import Panel from '../../components/Panel_default.svelte' 
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
<Loader loader_class="{loader_class}" loader_msg="{loader_msg}" />
<Panel
    on:handleRefresh={RefreshHalaman}
    panel_button_new={false}
    panel_button_refresh={true}
    panel_page="{page}"
    panel_total="{totalrecord}">
    <slot:template slot="panel_search">
        <div class="relative w-full">
            <div class="absolute inset-y-0 left-0 flex items-center pl-2">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 stroke-current text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
            </div>
            <input 
                bind:value={searchHome}
                type="text" placeholder="Search by Username, Page" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 focus:ring-0 focus:outline-none">
        </div>
    </slot:template>
    <slot:template slot="panel_body">
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
    </slot:template>
</Panel>



<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />



