<script>
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Modal_Alert from "../../components/Modal_alert.svelte"

    let page = "Admin Management";
    let sData = "New";
    let screen_height = screen.height - 480;
    let isModal_Form_New = true
    let isModalAlert = false
    let msg_error = "";
    let admin_listip = [];
    let tab_menu_1 = "bg-sky-600 text-white"
    let tab_menu_2 = ""
    let panel_edit = true
    let panel_iplist = false
    export let path_api = "";
    export let token = "";
    export let listHome = [];
    export let admin_listrule = [];
    export let totalrecord = 0;

    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        admin_username: yup
            .string()
            .required("Username is Required")
            .matches(
                /^[a-zA-z0-9]+$/,
                "Username must Character A-Z or a-z or 1-9"
            )
            .min(4, "Username must be at least 4 Character")
            .max(20, "Username must be at most 4 Character"),
        admin_password: yup
            .string()
            .required("Password is Required")
            .matches(
                /^[a-zA-z0-9]+$/,
                "Password must Character A-Z or a-z or 1-9"
            )
            .min(4, "Password must be at least 4 Character")
            .max(20, "Password must be at most 4 Character"),
        admin_name_field: yup.string().required("Name is Required"),
        admin_idrule_field: yup.number().required("Name is Required"),
        admin_status_field: yup.string().required("Status is Required"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            admin_username: "",
            admin_password: "",
            admin_name_field: "",
            admin_idrule_field: "0",
            admin_status_field: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(
                values.admin_username,
                values.admin_password,
                values.admin_name_field,
                values.admin_idrule_field
            );
        },
    });
    $: {
        if ($errors.admin_username || $errors.admin_password ||
            $errors.admin_name_field ||$errors.admin_idrule_field) {
            alert($errors.admin_username +
                    "\n" +
                    $errors.admin_password +
                    "\n" +
                    $errors.admin_name_field +
                    "\n" +
                    $errors.admin_idrule_field)
            $form.admin_username = "";
            $form.admin_password = "";
            $form.admin_name_field = "";
            $form.admin_idrule_field = "0";
            $form.admin_status_field = "";
        }
    }
    async function SaveTransaksi(username, password, name, rule) {
        let flag = true;
        msg_error = "";
        if (rule < 1) {
            flag = false;
            msg_error += "The Admin Rule is required";
        }
        if (flag) {
            const res = await fetch("/api/saveadmin", {
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
                    status: "ACTIVE",
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                $form.admin_username = "";
                $form.admin_password = "";
                $form.admin_name_field = "";
                $form.admin_idrule_field = "0";
            } else if (json.status == 403) {
                alert(json.message);
            } else {
            }
            RefreshHalaman();
        } else {
            alert(msg_error);
        }
    }
    
    async function EditData(e) {
        sData = "Edit"
        clearField();
        isModal_Form_New = true;
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
            logout();
        } else {
            for (let i = 0; i < record.length; i++) {
                $form.admin_username = e;
                $form.admin_password = "";
                $form.admin_name_field = record[i]["admin_nama"];
                $form.admin_idrule_field = parseInt(record[i]["admin_idrule"]);
                $form.admin_status_field = record[i]["admin_status"];
                // admin_create_field = record[i]["admin_create"];
                // admin_update_field = record[i]["admin_update"];
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
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const NewData = () => {
        sData = "New";
        clearField()
        isModal_Form_New = true;
    };
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
        $form.admin_username = "";
        $form.admin_password = "";
        $form.admin_name_field = "";
        $form.admin_idrule_field = "0";
        $form.admin_status_field = "";
    }
    let searchHome = "";
    let filterHome = [];
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.admin_username
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) ||
                    item.admin_nama
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
    console.log(admin_listrule)
</script>
<div class="flex flex-col">
    <div class="flex items-start">
      <h1 class=" text-black font-bold text-[2.25em] lg:text-[2.25em] md:text-[2.25em] sm:text-[2.25em]">Admin</h1>
      <div class="flex flex-1 justify-end items-center gap-2 mt-2">
        <button 
            on:click={() => {
                NewData();
            }}
            class="hidden btn btn-primary rounded-md lg:inline-flex">New</button>
        <button class="hidden btn btn-primary rounded-md lg:inline-flex">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>
      </div>
    </div>
    <div class="mt-2 flex gap-2">
      <div class="relative w-full">
        <div class="absolute inset-y-0 left-0 flex items-center pl-2">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 stroke-current text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
        <input 
          bind:value={searchHome}
          type="text" placeholder="Search" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4">
      </div>
    </div>
    <div class="mt-2 h-[650px]">
      <div class="overflow-x-auto">
        <table class="table table-compact w-full">
          <thead>
            <tr>
              <th width="1%">&nbsp;</th>
              <th width="1%" class="text-center text-[11px] lg:text-sm">NO</th> 
              <th width="1%" class="text-center text-[11px] lg:text-sm">STATUS</th>
              <th width="10%" class="text-left text-[11px] lg:text-sm">TIMEZONE</th> 
              <th width="10%" class="text-left text-[11px] lg:text-sm">IPADDRESS</th> 
              <th width="10%" class="text-center text-[11px] lg:text-sm">LAST LOGIN</th>
              <th width="10%" class="text-center text-[11px] lg:text-sm">JOIN DATE</th>
              <th width="10%" class="text-left text-[11px] lg:text-sm">RULE</th>
              <th width="15%" class="text-left text-[11px] lg:text-sm">USERNAME</th> 
              <th width="*" class="text-left text-[11px] lg:text-sm">NAMA</th>
            </tr>
          </thead> 
          <tbody>
            {#each filterHome as rec}
            <tr>
              <th class="cursor-pointer" on:click={() => {
                    EditData(rec.admin_username);
                }}>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                </svg>
              </th> 
              <th class="text-center  text-[11px] lg:text-sm">{rec.admin_no}</th> 
              <td  class="text-center  text-[11px] lg:text-sm">
                <div class="{rec.admin_statusclass} text-center rounded-md p-2 font-bold">{rec.admin_status}</div>
              </td> 
              <td class="text-left  text-[11px] lg:text-sm">{rec.admin_timezone}</td> 
              <td class="text-left text-[11px] lg:text-sm">{rec.admin_lastipaddres}</td> 
              <td class="text-center text-[11px] lg:text-sm">{rec.admin_lastlogin}</td> 
              <td class="text-center text-[11px] lg:text-sm">{rec.admin_joindate}</td> 
              <td class="text-left text-[11px] lg:text-sm">{rec.admin_rule}</td> 
              <td class="text-left text-[11px] lg:text-sm">{rec.admin_username}</td> 
              <td class="text-left text-[11px] lg:text-sm">{rec.admin_nama}</td> 
            </tr>
            {/each}
            
          </tbody> 
        </table>
      </div>
    </div>
    <div class="bg-gray-200 h-16 p-5">
      <span class="font-bold">TOTAL ROW : 100</span>
    </div>
</div>

<input type="checkbox" id="my-modal-formnew" class="modal-toggle" bind:checked={isModal_Form_New}>
<div class="modal" on:click|self={()=>isModal_Form_New = false}>
    <div class="modal-box relative select-none max-w-full lg:max-w-xl h-full lg:h-3/6 rounded-none lg:rounded-lg p-2  overflow-hidden">
        <div class="flex flex-col items-stretch">
            <div class="h-8">
                <label for="my-modal-formnew" class="btn btn-xs lg:btn-sm btn-circle absolute right-2 top-2">✕</label>
                <h3 class="text-xs lg:text-sm font-bold mt-1">Entry/{sData}</h3>
            </div>
            {#if sData == "Edit"}
                <div class="">
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
                </div>
            {/if}
            {#if panel_edit}
                <div class="flex flex-col h-[380px] overflow-auto gap-5 mt-2 ">
                    <div class="relative form-control mt-2">
                        <input
                            on:change="{handleChange}"
                            bind:value={$form.admin_username}
                            invalid={$errors.admin_username.length > 0}
                            type="text" 
                            id="username"
                            name="username"
                            placeholder="Username"
                            autocomplete="off"
                            class="peer w-full text-sm rounded px-3  border border-gray-300  focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none placeholder-transparent"> 
                            <label for="username" class="absolute left-3 top-[-0.7rem] text-gray-600 text-sm cursor-text 
                                transition-all
                                peer-placeholder-shown:text-base 
                                peer-placeholder-shown:text-gray-400 
                                peer-placeholder-shown:top-3 
                                peer-focus:text-[#1a73e8]
                                peer-focus:bg-[#fff]
                                peer-focus:text-[.75rem]
                                ">Username*</label>
                    </div>
                    <div class="relative form-control">
                        <input
                            on:change="{handleChange}"
                            bind:value={$form.admin_password}
                            invalid={$errors.admin_password.length > 0}
                            type="password" 
                            id="password"
                            name="password"
                            placeholder="Password"
                            autocomplete="off"
                            class="peer w-full rounded px-3  border border-gray-300  focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none placeholder-transparent"> 
                            <label for="password" class="absolute left-3 top-[-0.7rem] text-gray-600 text-sm cursor-text 
                                transition-all
                                peer-placeholder-shown:text-base 
                                peer-placeholder-shown:text-gray-400 
                                peer-placeholder-shown:top-3 
                                peer-focus:text-[#1a73e8]
                                peer-focus:bg-[#fff]
                                peer-focus:text-[.75rem]
                                ">Password*</label>
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
                    </div>
                    <div class="relative form-control">
                        <input
                            on:change="{handleChange}"
                            bind:value={$form.admin_name_field}
                            invalid={$errors.admin_name_field.length > 0}
                            type="text" 
                            id="name"
                            name="name"
                            placeholder="Name"
                            autocomplete="off"
                            class="peer w-full rounded px-3  border border-gray-300  focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none placeholder-transparent"> 
                            <label for="name" class="absolute left-3 top-[-0.7rem] text-gray-600 text-sm cursor-text 
                                transition-all
                                peer-placeholder-shown:text-base 
                                peer-placeholder-shown:text-gray-400 
                                peer-placeholder-shown:top-3 
                                peer-focus:text-[#1a73e8]
                                peer-focus:bg-[#fff]
                                peer-focus:text-[.75rem]
                                ">Name*</label>
                    </div>
                    {#if sData == "Edit"}
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
                    </div>
                    {/if}
                </div>
                <div class="h-20 ">
                    <div class="text-right align-middle  my-2">
                        <button
                            on:click={() => {
                                handleSubmit();
                            }}  
                            class="btn btn-primary ">Submit</button>
                    </div>
                </div>
            {/if}
            {#if panel_iplist}
                <div class="h-[380px] overflow-auto mt-2 ">
                    <table class="table table-compact w-full">
                        <thead>
                            <tr>
                                <th width="1%">&nbsp;</th>
                                <th width="1%">NO</th>
                                <th width="*">IPADDRESS</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each admin_listip as rec}
                            <tr>
                                <td class="cursor-pointer text-center">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                        <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                                    </svg>
                                </td>
                                <td class="text-center">{rec.adminiplist_no}</td>
                                <td class="text-left">{rec.adminiplist_iplist}</td>
                            </tr>
                            {/each}
                        </tbody>
                    </table>
                </div>
                <div class="h-20 ">
                    <div class="text-right align-middle  my-2">
                        <button
                            on:click={() => {
                                handleSubmit();
                            }}  
                            class="btn btn-primary ">New</button>
                    </div>
                </div>
            {/if}
        </div>
    </div>
</div>

<input type="checkbox" id="my-modal-alert" class="modal-toggle" bind:checked={isModalAlert}>
<Modal_Alert 
	modal_id="my-modal-alert" 
	modal_tipe="1" 
	modal_title="INFORMASI" 
	modal_title_class="" 
	modal_p_class="" 
	modal_widthheight_class="" 
	modal_message="{msg_error}" />