<script>
    import Home from "../admin/Home.svelte";
    export let path_api = ""
    let listHome = [];
    let record = "";
    let totalrecord = 0;
    let sData = "";
    let admin_username = "";
    let token = localStorage.getItem("token");
    let client_key = "";
    let akses_page = true;
    let admin_listrule = [];
    let admin_listip = [];
    let admin_name_field = "";
    let admin_type_field = "";
    let admin_idrule_field = "";
    let admin_status_field = "";
    let admin_create_field = "";
    let admin_update_field = "";
    async function initapp() {
        const res = await fetch(path_api+"api/home", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "ADMIN-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            // setTimeout(function(){ initPasaran() }, 1000);
            initAdmin();
        }
    }
    async function initAdmin() {
        const res = await fetch(path_api+"api/alladmin", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            let recordlistrule = json.listruleadmin;
            let status_class = "";
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    if (record[i]["admin_status"] == "ACTIVE") {
                        status_class = "bg-[#99FFCD]"
                    } else {
                        status_class = "bg-[#FF8080] text-white"
                    }
                    listHome = [
                        ...listHome,
                        {
                            admin_no: record[i]["admin_no"],
                            admin_username: record[i]["admin_username"],
                            admin_nama: record[i]["admin_nama"],
                            admin_tipe: record[i]["admin_tipe"],
                            admin_rule: record[i]["admin_rule"],
                            admin_timezone: record[i]["admin_timezone"],
                            admin_joindate: record[i]["admin_joindate"],
                            admin_lastlogin: record[i]["admin_lastlogin"],
                            admin_lastipaddres: record[i]["admin_lastipaddres"],
                            admin_status: record[i]["admin_status"],
                            admin_statusclass: status_class,
                        },
                    ];
                }
            }
            if (recordlistrule != null) {
                for (let i = 0; i < recordlistrule.length; i++) {
                    admin_listrule = [
                        ...admin_listrule,
                        {
                            adminrule_idruleadmin:
                                recordlistrule[i]["adminrule_idruleadmin"],
                            adminrule_name: recordlistrule[i]["adminrule_name"],
                        },
                    ];
                }
            }
        } else {
            logout();
        }
    }
    async function editAdmin(e) {
        clearEdit();
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
                admin_name_field = record[i]["admin_nama"];
                admin_idrule_field = parseInt(record[i]["admin_idrule"]);
                admin_status_field = record[i]["admin_status"];
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
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        admin_listrule = [];
        totalrecord = 0;
        setTimeout(function () {
            initAdmin();
        }, 1000);
    };
    const handleEditData = (e) => {
        admin_username = e.detail.e;
        admin_type_field = e.detail.f;
        sData = "Edit";
        editAdmin(admin_username);
    };
    const handleLogout = (e) => {
        logout()
    }
    function clearEdit(){
        admin_name_field = "";
        admin_idrule_field = "";
        admin_status_field = "";
        admin_create_field = "";
        admin_update_field = "";
    }
    initapp();
</script>

<Home
    on:handleRefreshData={handleRefreshData}
    on:handleEditData={handleEditData}
    on:handleLogout={handleLogout}
    {path_api}
    {token}
    {admin_listrule}
    {listHome}
    {totalrecord}/>