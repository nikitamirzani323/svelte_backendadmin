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
    export let admin_listrule = [];
    export let totalrecord = 0;

    let page = "Admin Management";
    let sData = "New";
    let isModal_Form_New = false
    let isModal_Form_Listipaddress = false
    let isModalLoading = false
    let isModalNotif = false
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let msg_error = "";
    let admin_listip = [];
    let tab_menu_1 = "bg-sky-600 text-white"
    let tab_menu_2 = ""
    let panel_edit = true
    let panel_iplist = false
    let admin_tipe = "ADMIN";
    let searchHome = "";
    let filterHome = [];
    let form_field_ipaddress = "";
    let isInput_username_enabled = true;
    let admin_create_field = ""
    let admin_update_field = ""
    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        admin_username_field: yup
            .string()
            .required("Username is Required")
            .matches(
                /^[a-zA-z0-9]+$/,
                "Username must Character A-Z or a-z or 1-9"
            )
            .min(4, "Username must be at least 4 Character")
            .max(20, "Username must be at most 20 Character"),
        admin_password_field: yup.string(),
        admin_name_field: yup
            .string()
            .required("Nama is Required")
            .matches(
                /^[a-zA-z0-9]+$/,
                "Nama must Character A-Z or a-z or 1-9"
            )
            .min(4, "Nama must be at least 4 Character")
            .max(20, "Nama must be at most 20 Character"),
        admin_idrule_field: yup.number().required("Admin Rule is Required"),
        admin_status_field: yup.string().required("Status is Required"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            admin_username_field: "",
            admin_password_field: "",
            admin_name_field: "",
            admin_idrule_field: "0",
            admin_status_field: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(
                values.admin_username_field,
                values.admin_password_field,
                values.admin_name_field,
                values.admin_idrule_field,
                values.admin_status_field
            );
        },
    });
    async function SaveTransaksi(username, password, name, rule,status) {
        let flag = true;
        msg_error = "";
        const regexExp = /^[a-zA-z0-9]+$/gi;
        let flag_password = regexExp.test(password)
        if(password != ""){
            if(!flag_password){
                flag = false;
                msg_error += "The Format Password invalid\n Password must Character A-Z or a-z or 1-9";
            }
        }
        if (rule == "0") {
            flag = false;
            msg_error += "The Admin Rule is required";
        }
        if(status == ""){
            flag = false;
            msg_error += "The Status is required";
        }
        if (flag) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/saveadmin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    idruleadmin: parseInt(rule),
                    page: "ADMIN-SAVE",
                    username: username,
                    password: password,
                    nama: name,
                    status: status,
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
                        $form.admin_username_field = "";
                        $form.admin_password_field = "";
                        $form.admin_name_field = "";
                        $form.admin_idrule_field = "0";
                        $form.admin_status_field = "";
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
    async function SaveIpaddress() {
        let flag = true;
        let totaliplist = admin_listip.length;
        msg_error = "";
        if (form_field_ipaddress == "") {
            flag = false;
            msg_error += "The Ipaddress is required\n";
        }
        const regexExp = /^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/gi;
        let flag_ip = regexExp.test(form_field_ipaddress)
        if(!flag_ip){
            flag = false;
            msg_error += "The Format Ipaddress invalid\n";
        }
        if(totaliplist > 5){
            flag = false;
            msg_error += "Maximal 5 Ipaddress Active\n";
        }
        if (flag) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/saveadminiplist", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sData: "New",
                    page: "ADMIN-SAVE",
                    username: $form.admin_username_field,
                    ipaddress: form_field_ipaddress,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                loader_msg = json.message
                EditData($form.admin_username_field)
                form_field_ipaddress = "";
            } else if (json.status == 403) {
                loader_msg = json.message
            } else {
                loader_msg = json.message;
            }
            buttonLoading_class = "btn btn-primary"
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
        } else {
            alert(msg_error);
        }
    }
    async function deleteIpList(e) {
        const res = await fetch(path_api+"api/deleteadminiplist", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idcompiplist: parseInt(e),
                username: $form.admin_username_field,
                page:"ADMIN-SAVE",
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            admin_listip = [];
            EditData($form.admin_username_field)
        }else if(json.status == 403){
            alert(json.message)
        }
    }
    async function EditData(e,x) {
        if(e != ""){
            admin_tipe = x;
            sData = "Edit";
            clearField();
            isModalLoading = true;
            isInput_username_enabled = false;
            
            admin_create_field = "";
            admin_update_field = "";
            const res = await fetch(path_api+"api/editadmin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    username: e,
                }),
            });
            const json = await res.json();
            let record = json.record;
            let recordlistrule = json.listruleadmin;
            let recordlistip = json.listip;
            if (json.status === 400) {
                dispatch("handleLogout", "call");
            }else if(json.status === 200){
                for (let i = 0; i < record.length; i++) {
                    $form.admin_username_field = e;
                    $form.admin_password_field = "";
                    $form.admin_name_field = record[i]["admin_nama"];
                    $form.admin_idrule_field = record[i]["admin_idrule"];
                    $form.admin_status_field = record[i]["admin_status"];
                    admin_create_field = record[i]["admin_create"];
                    admin_update_field = record[i]["admin_update"];
                }
                if (recordlistrule != null) {
                    for (let i = 0; i < recordlistrule.length; i++) {
                        admin_listrule = [
                            ...admin_listrule,
                            {
                                adminrule_idruleadmin:recordlistrule[i]["adminrule_idruleadmin"],
                                adminrule_name: recordlistrule[i]["adminrule_name"],
                            },
                        ];
                    }
                }
                if (recordlistip != null) {
                    let no = 0;
                    for (let i = 0; i < recordlistip.length; i++) {
                        no = no + 1;
                        admin_listip = [
                            ...admin_listip,
                            {
                                adminiplist_no: no,
                                adminiplist_idcompiplist:recordlistip[i]["adminiplist_idcompiplist"],
                                adminiplist_iplist:recordlistip[i]["adminiplist_iplist"],
                            },
                        ];
                    }
                }
                isModalLoading = false;
                isModal_Form_New = true;
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
        isInput_username_enabled = true;
        clearField()
        isModal_Form_New = true;
    };
    const handleNewListIp = () => {
        isModal_Form_Listipaddress = true;
    }
    const ChangeTabMenu = (e) => {
        if(e == "menu_2"){
            tab_menu_1 = ""
            tab_menu_2 = "bg-sky-600 text-white"
            panel_edit = false
            panel_iplist = true
        }else{
            tab_menu_1 = "bg-sky-600 text-white"
            tab_menu_2 = ""
            panel_edit = true
            panel_iplist = false
        }
    }
    function clearField(){
        if(sData == "Edit"){
            admin_listrule = []
            admin_listip = []
        } 
        $form.admin_username_field = "";
        $form.admin_password_field = "";
        $form.admin_name_field = "";
        $form.admin_idrule_field = "0";
        $form.admin_status_field = "";
        form_field_ipaddress = "";
    }
    
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.admin_username
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) ||
                    item.admin_nama
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.admin_rule
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
            type="text" placeholder="Search by Rule, Username, Nama" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 focus:ring-0 focus:outline-none">
    </slot:template>
    <slot:template slot="panel_body">
        <table class="table table-compact w-full">
            <thead class="sticky top-0">
                <tr>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center"></th>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">NO</th>
                    <th width="5%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">STATUS</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">TIMEZONE</th>
                    <th width="15%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">IPADDRESS</th>
                    <th width="15%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">LAST LOGIN</th>
                    <th width="15%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">JOIN DATE</th>
                    <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">RULE</th>
                    <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">USERNAME</th>
                    <th width="*" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">NAMA</th>
                </tr>
            </thead>
            {#if filterHome != ""}
                <tbody>
                    {#each filterHome as rec}
                    <tr>
                        <td on:click={() => {
                            EditData(rec.admin_username,rec.admin_tipe);
                            }} class="text-center text-xs cursor-pointer">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                            </svg>
                        </td>
                        <td class="text-xs lg:text-sm align-top text-center">{rec.admin_no}</td>
                        <td class="text-xs lg:text-sm align-top text-center"><span class="{rec.admin_statusclass} text-center rounded-md p-1 px-2 shadow-lg">{rec.admin_status}</span></td>
                        <td class="text-xs lg:text-sm align-top text-center">{rec.admin_timezone}</td>
                        <td class="text-xs lg:text-sm align-top text-center">{rec.admin_lastipaddres}</td>
                        <td class="text-xs lg:text-sm align-top text-center">{rec.admin_lastlogin}</td>
                        <td class="text-xs lg:text-sm align-top text-center">{rec.admin_joindate}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.admin_rule}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.admin_username}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.admin_nama}</td>
                    </tr>
                    {/each}
                </tbody>
            {:else}
                <tbody>
                    <tr>
                        <td colspan="10" class="text-center">
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
    modal_popup_class="select-none max-w-full lg:max-w-xl overflow-hidden">
    <slot:template slot="modalpopup_body">
        {#if sData == "Edit" && admin_tipe == "ADMIN"}
            <ul class="flex justify-center items-center gap-2">
                <li on:click={() => {
                        ChangeTabMenu("menu_1");
                    }}
                    class="items-center {tab_menu_1}  px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">Edit</li>
                <li on:click={() => {
                        ChangeTabMenu("menu_2");
                    }}
                    class="items-center {tab_menu_2} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">List Ipaddress</li>
            </ul>
        {/if}
        {#if panel_edit}
            <div class="flex flex-auto flex-col overflow-auto gap-5 mt-2 ">
                <div class="relative form-control mt-2">
                    <Input_custom
                        input_onchange="{handleChange}"
                        input_autofocus={false}
                        input_required={true}
                        input_tipe="text"
                        input_invalid={$errors.admin_username_field.length > 0}
                        bind:value={$form.admin_username_field}
                        input_id="admin_username_field"
                        input_enabled={isInput_username_enabled}
                        input_placeholder="Username"/>
                    {#if $errors.admin_username_field}
                        <small class="text-pink-600 text-[11px]">{$errors.admin_username_field}</small>
                    {/if}
                </div>
                <div class="relative form-control">
                    <Input_custom
                        input_onchange="{handleChange}"
                        input_autofocus={false}
                        input_required={true}
                        input_tipe="password"
                        input_attr="password"
                        input_invalid={$errors.admin_password_field.length > 0}
                        bind:value={$form.admin_password_field}
                        input_id="admin_password_field"
                        input_placeholder="Password"/>
                    {#if $errors.admin_password_field}
                        <small class="text-pink-600 text-[11px]">{$errors.admin_password_field}</small>
                    {/if}
                </div>
                <div class="relative form-control">
                    <select
                        on:change="{handleChange}"
                        bind:value={$form.admin_idrule_field}
                        invalid={$errors.admin_idrule_field.length > 0} 
                        class="w-full rounded px-3  border border-gray-300 focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none">
                        <option disabled selected value="0">--Pilih Admin Rule--</option>
                        {#each admin_listrule as rec}
                        <option value="{rec.adminrule_idruleadmin}">{rec.adminrule_name}</option>
                        {/each}
                    </select>
                    {#if $errors.admin_idrule_field}
                        <small class="text-pink-600 text-[11px]">{$errors.admin_idrule_field}</small>
                    {/if}
                </div>
                <div class="relative form-control">
                    <Input_custom
                        input_onchange="{handleChange}"
                        input_autofocus={false}
                        input_required={true}
                        input_tipe="text"
                        input_invalid={$errors.admin_name_field.length > 0}
                        bind:value={$form.admin_name_field}
                        input_id="admin_name_field"
                        input_placeholder="Nama"/>
                    {#if $errors.admin_name_field}
                        <small class="text-pink-600 text-[11px]">{$errors.admin_name_field}</small>
                    {/if}
                </div>
                <div class="relative form-control">
                    <select
                        on:change="{handleChange}"
                        bind:value={$form.admin_status_field}
                        invalid={$errors.admin_status_field.length > 0} 
                        class="w-full rounded px-3  border border-gray-300 focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none">
                        <option disabled selected value="">--Pilih Status--</option>
                        <option value="ACTIVE">ACTIVE</option>
                        <option value="BANNED">BANNED</option>
                    </select>
                    {#if $errors.admin_status_field}
                        <small class="text-pink-600 text-[11px]">{$errors.admin_status_field}</small>
                    {/if}
                </div>
                {#if sData == "Edit"}
                <div class="text-[11px]">
                    Create : {admin_create_field} <br>
                    Update : {admin_update_field}
                </div>
                {/if}
            </div>
            <div class="flex flex-wrap justify-end align-middle p-[0.75rem] mt-2">
                <button
                    on:click={() => {
                        handleSubmit();
                    }}  
                    class="{buttonLoading_class}">Submit</button>
            </div>
        {/if}
        {#if panel_iplist}
            <div class="flex-auto h-[380px] overflow-auto mt-2 ">
                <table class="table table-compact w-full">
                    <thead>
                        <tr>
                            <th width="1%">&nbsp;</th>
                            <th width="1%">NO</th>
                            <th width="*">IPADDRESS</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#if admin_listip != ""}
                            {#each admin_listip as rec}
                            <tr>
                                <td on:click={() => {
                                    deleteIpList(
                                        rec.adminiplist_idcompiplist
                                    );
                                    }} class="cursor-pointer text-center">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                        <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                                    </svg>
                                </td>
                                <td class="text-center">{rec.adminiplist_no}</td>
                                <td class="text-left">{rec.adminiplist_iplist}</td>
                            </tr>
                            {/each}
                        {:else}
                            <tr>
                                <td colspan="3" class="text-left text-xs font-semibold">No Record</td>
                            </tr>
                        {/if}
                    </tbody>
                </table>
            </div>
            <div class="flex flex-wrap justify-end align-middle p-[0.75rem] mt-2">
                <button
                    on:click={() => {
                        handleNewListIp();
                    }}  
                    class="btn btn-primary ">New</button>
            </div>
        {/if}
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-formipaddress" class="modal-toggle" bind:checked={isModal_Form_Listipaddress}>
<Modal_popup
    modal_popup_id="my-modal-formipaddress"
    modal_popup_title="New IPAddress"
    modal_popup_class="select-none max-w-full lg:max-w-xl overflow-hidden">
    <slot:template slot="modalpopup_body">
        <div class="flex flex-auto flex-col overflow-auto gap-5 mt-2 ">
            <div class="mt-2">
                <Input_custom
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="text"
                    bind:value={form_field_ipaddress}
                    input_id="ipaddress"
                    input_placeholder="IPAddress"/>
            </div>
        </div>
        <div class="flex flex-wrap justify-end align-middle p-[0.75rem] mt-2">
            <button
                on:click={() => {
                    SaveIpaddress();
                }}  
                class="{buttonLoading_class}">Submit</button>
        </div>
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />



