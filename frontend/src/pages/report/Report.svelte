<script>
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Input_custom from '../../components/Input.svelte' 
    
    import dayjs from "dayjs";
    export let path_api = ""
    let listwinlose = [];
    let startdate = ""
    let enddate = ""
    let token = localStorage.getItem("token");
    let akses_page = false;
    let isModalNotif = false;
    let total_turnover = 0;
    let total_winlose = 0;
    let total_winlose_agent = 0;
    let css_sub_winlosemember = "text-blue-700 font-semibold";
    let css_sub_winloseagent = "text-blue-700 font-semibold";
    let msg_error = ""
    async function initapp() {
        const res = await fetch(path_api+"api/home", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "REPORTWINLOSE-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            akses_page = true;
        }
    }
    
    async function call_generatereport() {
        listwinlose = [];
        total_turnover = 0;
        total_winlose = 0;
        total_winlose_agent = 0;
        let flag = false;
        let date1 = dayjs(startdate);
        let date2 = dayjs(enddate);
        const diff = date2.diff(date1, "day", true);
        const days = Math.floor(diff);
        if (days > 0 && days < 31) {
            flag = true;
        }
        if (flag) {
            const res = await fetch(path_api+"api/reportwinlose", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    client_start: startdate,
                    client_end: enddate,
                }),
            });
            const json = await res.json();
            let record = json.record;
            if (json.status == 200) {
                if (record != null) {
                    for (var i = 0; i < record.length; i++) {
                        let css_winlosemember = "text-red-500 font-semibold";
                        let css_winloseagent = "text-red-500 font-semibold";
                        if (record[i]["report_client_winlose"] > 0) {
                            css_winlosemember = "text-blue-700 font-semibold";
                        }
                        if (record[i]["report_agent_winlose"] > 0) {
                            css_winloseagent = "text-blue-700 font-semibold";
                        }
                        listwinlose = [
                            ...listwinlose,
                            {
                                report_client_username:record[i]["report_client_username"],
                                report_client_turnover:record[i]["report_client_turnover"],
                                report_client_winlose:record[i]["report_client_winlose"],
                                report_client_winlose_css: css_winlosemember,
                                report_agent_winlose:record[i]["report_agent_winlose"],
                                report_agent_winlose_css: css_winloseagent,
                            },
                        ];
                    }
                    total_turnover = parseInt(json.subtotalturnover);
                    total_winlose = parseInt(json.subtotalwinlose);
                    total_winlose_agent = parseInt(json.subtotalwinlosecompany);
                    if (total_winlose > 0) {
                        css_sub_winlosemember = "text-blue-700 font-semibold";
                    }else{
                        css_sub_winlosemember = "text-red-500 font-semibold";
                    }
                    if (total_winlose_agent > 0) {
                        css_sub_winloseagent = "text-blue-700 font-semibold";
                    }else{
                        css_sub_winloseagent = "text-red-500 font-semibold";
                    }
                    
                }
            } else if (json.status == 400) {
                logout();
            }
        } else {
            msg_error ="Report Winlose, hanya bisa dilakukan dengan range tanggal 31 hari";
            isModalNotif = true;
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    initapp();
</script>
{#if akses_page == true}
    <div class=" mx-auto px-2  ">
        <div class="flex gap-2">
            <div class="bg-white shadow-lg  w-[500px] ">
                <div class="flex flex-col w-full">
                    <div class="border-slate-200">
                        <div class="flex justify-start items-center border-b-2 p-2">
                            <h1 class="text-slate-600 font-bold text-sm lg:text-xl uppercase w-full">Report Winlose</h1>
                            <div class="flex justify-end w-full p-2">
                                <button on:click={() => {
                                    call_generatereport();
                                }} class="btn btn-primary hover:bg-primary m-0 p-2 h-1 min-h-[40px]  rounded-md shadow-lg shadow-[#b3e4fc] border-none">Check</button>
                            </div>
                        </div>
                    </div>
                    <div class="w-full mt-2 p-2 flex flex-col gap-5 ">
                        <Input_custom
                            input_required={true}
                            input_tipe="date"
                            bind:value={startdate}
                            input_id="startdate"
                            input_placeholder="Start"/>
                        
                        <Input_custom
                            input_required={true}
                            input_tipe="date"
                            bind:value={enddate}
                            input_id="enddate"
                            input_placeholder="End"/>
                    </div>
                </div>
            </div>
            <div class="bg-white shadow-lg p-2 w-full">
                <div class="flex flex-col w-full">
                    <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[550px] overflow-y-scroll">
                        <table class="table table-compact w-full ">
                            <thead class="sticky top-0">
                                <tr>
                                    <th width="*" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">USERNAME</th>
                                    <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">TURNOVER</th>
                                    <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">WINLOSE MEMBER</th>
                                    <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">WINLOSE AGENT</th>
                                </tr>
                            </thead>
                            {#if listwinlose != ""}
                                <tbody>
                                    {#each listwinlose as rec}
                                    <tr>
                                        <td class="text-xs lg:text-sm align-top text-left whitespace-nowrap">{rec.report_client_username}</td>
                                        <td class="text-xs lg:text-sm align-top text-right text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.report_client_turnover)}</td>
                                        <td class="text-xs lg:text-sm align-top text-right {rec.report_client_winlose_css} whitespace-nowrap">{new Intl.NumberFormat().format(rec.report_client_winlose)}</td>
                                        <td class="text-xs lg:text-sm align-top text-right {rec.report_agent_winlose_css} whitespace-nowrap">{new Intl.NumberFormat().format(rec.report_agent_winlose)}</td>
                                    </tr>
                                    {/each}
                                </tbody>
                            {:else}
                                <tbody>
                                    <tr>
                                        <td colspan="4" class="text-center">
                                            <progress class="self-start progress progress-primary w-56"></progress>
                                        </td>
                                    </tr>
                                </tbody>
                            {/if}
                        </table>
                    </div>
                    <div class="bg-[#F7F7F7] rounded-sm h-20 p-2 ">
                        <table class="w-full">
                            <tr>
                                <td class="text-left font-semibold text-xs">TOTAL TURNOVER</td>
                                <td class="text-right text-xs text-blue-700 font-semibold">{new Intl.NumberFormat().format(total_turnover)}</td>
                            </tr>
                            <tr>
                                <td class="text-left font-semibold text-xs">MEMBER WINLOSE</td>
                                <td class="text-right text-xs {css_sub_winlosemember}">{new Intl.NumberFormat().format(total_winlose)}</td>
                            </tr>
                            <tr>
                                <td class="text-left font-semibold text-xs">COMPANY WINLOSE</td>
                                <td class="text-right text-xs {css_sub_winloseagent}">{new Intl.NumberFormat().format(total_winlose_agent)}</td>
                            </tr>
                        </table>
                    </div>
                </div>
            </div>
        </div>
        
    </div>
{/if}

<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />