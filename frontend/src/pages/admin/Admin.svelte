<script>
    import Home from "../admin/Home.svelte";
    export let path_api = ""
    let listHome = [];
    let record = "";
    let totalrecord = 0;
    let token = localStorage.getItem("token");
    let akses_page = false;
    let admin_listrule = [];
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
            akses_page = true;
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
                        status_class = "bg-[#8BC34A] "
                    } else {
                        status_class = "bg-[#E91E63] text-white"
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
    const handleLogout = (e) => {
        logout()
    }
    initapp();
</script>
{#if akses_page == true}
    <Home
        on:handleRefreshData={handleRefreshData}
        on:handleLogout={handleLogout}
        {path_api}
        {token}
        {admin_listrule}
        {listHome}
        {totalrecord}/>
{/if}