<script>
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Input_custom from '../../components/Input.svelte' 
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Modal_popup from '../../components/Modal_popup.svelte' 
    import Loader from '../../components/Loader.svelte' 
    import Panel from '../../components/Panel_default.svelte' 

    export let path_api = "";
    export let token = "";
    export let listHome = [];
    export let totalrecord = 0;

    let page = "Admin Rule";
    let sData = "New";
    let isModal_Form_New = false
    let isModalLoading = false
    let isModalNotif = false
    let modal_width = "max-w-xl"
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let buttonLoading2_class = "btn btn-primary"
    let msg_error = "";
    let adminrule_id = "";
    let adminrule_rule_field = "";
    let searchHome = "";
    let filterHome = [];

    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        home_name_field: yup
            .string()
            .required("Rule is Required")
            .matches(
                /^[a-zA-z0-9 ]+$/,
                "Rule must Character A-Z or a-z or 1-9"
            )
            .min(4, "Rule must be at least 4 Character")
            .max(50, "Rule must be at most 50 Character"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            home_name_field: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(values.home_name_field);
        },
    });
    async function SaveTransaksi(name) {
        let flag = true;
        msg_error = "";
        
        if (flag) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/saveadminrule", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page: "ADMINRULE-SAVE",
                    idrule: adminrule_id,
                    nama: name,
                }),
            });
            const json = await res.json();
            if(!res.ok){
                loader_msg = "System Mengalami Trouble"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
            }else{
                if (json.status == 200) {
                    loader_msg = json.message
                    if(sData == "New"){
                        $form.home_name_field = "";
                    }
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_class = "btn btn-primary"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
                RefreshHalaman();
            }
        } else {
            alert(msg_error);
        }
    }
    async function Updateconfig() {
        let flag = false;
        msg_error = "";
        if (adminrule_rule_field == "") {
            flag = true;
            msg_error += "The List is required\n";
        }
        if (flag == false) {
            buttonLoading2_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/saveadminruleconf", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page: "ADMINRULE-SAVE",
                    idrule: adminrule_id,
                    rule: adminrule_rule_field.toString(),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                loader_msg = json.message;
            } else if (json.status == 403) {
                loader_msg = json.message;;
            } else {
                loader_msg = json.message;
            }
            buttonLoading2_class = "btn btn-primary"
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
        } else {
            alert(msg_error);
        }
    }
    async function EditData(e,y) {
        if(e != ""){
            adminrule_id = e;
            adminrule_rule_field = "";
            isModalLoading = true;
            modal_width = "max-w-4xl";
            sData = "Edit";
            $form.home_name_field = y;
            const res = await fetch(path_api+"api/editadminrule", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idrule: e,
                }),
            });
            const json = await res.json();
            let record = json.record;
            if (json.status === 400) {
                dispatch("handleLogout", "call");
            }else if(json.status === 200) {
                isModal_Form_New = true;
                isModalLoading = false;
                for (let i = 0; i < record.length; i++) {
                    adminrule_rule_field = record[i]["adminrule_rule"];
                }
            }else{
                isModalLoading = false;
                isModalNotif = true;
                msg_error = "Silahkan Hubungi Administrator"
            }
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const NewData = () => {
        sData = "New";
        clearField()
        adminrule_id = "";
        modal_width = "max-w-xl"
        isModal_Form_New = true;
    };
    
   
    function clearField(){
        $form.home_name_field = "";
    }
    
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_nama
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
    on:handleNewForm={NewData}
    on:handleRefresh={RefreshHalaman}
    panel_button_new={true}
    panel_button_refresh={true}
    panel_page="{page}"
    panel_total="{totalrecord}">
    <slot:template slot="panel_search">
        <div class="absolute inset-y-0 left-0 flex items-center pl-2">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 stroke-current text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
        </div>
        <input 
            bind:value={searchHome}
            type="text" placeholder="Search by Rule" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 focus:ring-0 focus:outline-none">
    </slot:template>
    <slot:template slot="panel_body">
        <table class="table table-compact w-full ">
            <thead class="sticky top-0">
                <tr>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center"></th>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">NO</th>
                    <th width="*" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">RULE</th>
                </tr>
            </thead>
            {#if filterHome != ""}
                <tbody>
                    {#each filterHome as rec}
                    <tr>
                        <td on:click={() => {
                            EditData(rec.home_id,rec.home_nama);
                            }} class="text-center text-xs cursor-pointer">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                            </svg>
                        </td>
                        <td class="text-xs lg:text-sm align-top text-center">{rec.home_no}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_nama}</td>
                    </tr>
                    {/each}
                </tbody>
            {:else}
                <tbody>
                    <tr>
                        <td colspan="3" class="text-center">
                            <progress class="self-start progress progress-primary w-56"></progress>
                        </td>
                    </tr>
                </tbody>
            {/if}
        </table>
    </slot:template>
</Panel>


<input type="checkbox" id="my-modal-formnew" class="modal-toggle" bind:checked={isModal_Form_New}>
<Modal_popup
    modal_popup_id="my-modal-formnew"
    modal_popup_title="Entry/{sData}"
    modal_popup_class="select-none w-11/12 {modal_width} overflow-hidden">
    <slot:template slot="modalpopup_body">
        {#if sData=="New"}
            <div class="flex flex-auto flex-col overflow-auto gap-5 mt-2 ">
                <div class="relative form-control mt-2">
                    <Input_custom
                        input_onchange="{handleChange}"
                        input_autofocus={false}
                        input_required={true}
                        input_tipe="text"
                        input_invalid={$errors.home_name_field.length > 0}
                        bind:value={$form.home_name_field}
                        input_id="home_name_field"
                        input_placeholder="Rule"/>
                    {#if $errors.home_name_field}
                        <small class="text-pink-600 text-[11px]">{$errors.home_name_field}</small>
                    {/if}
                </div>
            </div>
            <div class="flex flex-wrap justify-end align-middle p-[0.75rem] mt-2">
                <button
                    on:click={() => {
                        handleSubmit();
                    }}  
                    class="{buttonLoading_class}">Submit</button>
            </div>
        {/if}
        {#if sData=="Edit"}
            <div class="flex justify-between  gap-2">
                <div class="w-full">
                    <div class="flex flex-auto flex-col overflow-auto gap-5 mt-2  ">
                        <div class="relative form-control mt-2">
                            <Input_custom
                                input_onchange="{handleChange}"
                                input_autofocus={false}
                                input_required={true}
                                input_tipe="text"
                                input_invalid={$errors.home_name_field.length > 0}
                                bind:value={$form.home_name_field}
                                input_id="home_name_field"
                                input_placeholder="Rule"/>
                            {#if $errors.home_name_field}
                                <small class="text-pink-600 text-[11px]">{$errors.home_name_field}</small>
                            {/if}
                        </div>
                    </div>
                    <div class="flex flex-wrap justify-end align-middle  mt-2">
                        <button
                            on:click={() => {
                                handleSubmit();
                            }}  
                            class="{buttonLoading_class} btn-block">Submit</button>
                    </div>
                </div>
                <div class="w-full p-2">
                    <h2 class="text-lg font-bold mb-2">Setting Admin Rule</h2>
                    <table class="table table-compact w-full">
                        <thead>
                            <tr>
                                <th colspan="2">DASHBOARD</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td width="1%">
                                    <input
                                    bind:group={adminrule_rule_field}
                                    type="checkbox"
                                    value="DASHBOARD-VIEW"/>
                                </td>
                                <td width="*">VIEW</td>
                            </tr>
                        </tbody>
                    </table>
                    <table class="table table-compact w-full mt-1">
                        <thead>
                            <tr>
                                <th colspan="2">PERIODE</th>
                                <th colspan="2">PREDIKSI</th>
                                <th colspan="2">REPORT WINLOSE</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="PERIODE-VIEW"/>
                                </td>
                                <td width="*">VIEW</td>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="PREDIKSI-VIEW"/>
                                </td>
                                <td width="*">VIEW</td>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="REPORTWINLOSE-VIEW"/>
                                </td>
                                <td width="*">VIEW</td>
                            </tr>
                            <tr>
                                <td width="1%">
                                    <input
                                    bind:group={adminrule_rule_field}
                                    type="checkbox"
                                    value="PERIODE-SAVE"/>
                                </td>
                                <td width="*">SAVE</td>
                                <td colspan="4">&nbsp;</td>
                            </tr>
                        </tbody>
                    </table>
                    <table class="table table-compact w-full mt-1">
                        <thead>
                            <tr>
                                <th colspan="2">PASARAN</th>
                                <th colspan="2">ADMIN MANAGEMENT</th>
                                <th colspan="2">ADMIN RULE</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="PASARAN-VIEW"/>
                                </td>
                                <td width="*">VIEW</td>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="ADMIN-VIEW"/>
                                </td>
                                <td width="*">VIEW</td>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="ADMINRULE-VIEW"/>
                                </td>
                                <td width="*">VIEW</td>
                            </tr>
                            <tr>
                                <td width="1%">
                                    <input
                                    bind:group={adminrule_rule_field}
                                    type="checkbox"
                                    value="PASARAN-SAVE"/>
                                </td>
                                <td width="*">SAVE</td>
                                <td width="1%">
                                    <input
                                    bind:group={adminrule_rule_field}
                                    type="checkbox"
                                    value="ADMIN-SAVE"/>
                                </td>
                                <td width="*">SAVE</td>
                                <td width="1%">
                                    <input
                                    bind:group={adminrule_rule_field}
                                    type="checkbox"
                                    value="ADMINRULE-SAVE"/>
                                </td>
                                <td width="*">SAVE</td>
                            </tr>
                        </tbody>
                    </table>  
                    <button
                        on:click={() => {
                            Updateconfig();
                        }}  
                        class="{buttonLoading2_class} btn-block">Submit</button> 
                </div>
            </div>
        {/if}
    </slot:template>
</Modal_popup>


<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />
