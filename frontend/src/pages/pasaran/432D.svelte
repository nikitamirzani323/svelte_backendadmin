<script>
    import { createEventDispatcher } from "svelte";
    import Input_custom from '../../components/Input.svelte'

    export let pasaran_tipe = "";
    export let path_api = "";
    export let token = "";
    export let idcomppasaran = "";
    export let pasaran_idpasarantogel_field = "";
    export let pasaran_minbet_432d_field = 0;
    export let pasaran_maxbet4d_432d_field = 0;
    export let pasaran_maxbet3d_432d_field = 0;
    export let pasaran_maxbet3dd_432d_field = 0;
    export let pasaran_maxbet2d_432d_field = 0;
    export let pasaran_maxbet2dd_432d_field = 0;
    export let pasaran_maxbet2dt_432d_field = 0;
    export let pasaran_limitotal4d_432d_field = 0;
    export let pasaran_limitotal3d_432d_field = 0;
    export let pasaran_limitotal3dd_432d_field = 0;
    export let pasaran_limitotal2d_432d_field = 0;
    export let pasaran_limitotal2dd_432d_field = 0;
    export let pasaran_limitotal2dt_432d_field = 0;
    export let pasaran_limitglobal4d_432d_field = 0;
    export let pasaran_limitglobal3d_432d_field = 0;
    export let pasaran_limitglobal3dd_432d_field = 0;
    export let pasaran_limitglobal2d_432d_field = 0;
    export let pasaran_limitglobal2dd_432d_field = 0;
    export let pasaran_limitglobal2dt_432d_field = 0;
    export let pasaran_disc4d_432d_field = 0;
    export let pasaran_disc3d_432d_field = 0;
    export let pasaran_disc3dd_432d_field = 0;
    export let pasaran_disc2d_432d_field = 0;
    export let pasaran_disc2dd_432d_field = 0;
    export let pasaran_disc2dt_432d_field = 0;
    export let pasaran_win4d_432d_field = 0;
    export let pasaran_win3d_432d_field = 0;
    export let pasaran_win3dd_432d_field = 0;
    export let pasaran_win2d_432d_field = 0;
    export let pasaran_win2dd_432d_field = 0;
    export let pasaran_win2dt_432d_field = 0;
    export let pasaran_win4dnodisc_432d_field = 0;
    export let pasaran_win3dnodisc_432d_field = 0;
    export let pasaran_win3ddnodisc_432d_field = 0;
    export let pasaran_win2dnodisc_432d_field = 0;
    export let pasaran_win2ddnodisc_432d_field = 0;
    export let pasaran_win2dtnodisc_432d_field = 0;
    export let pasaran_win4d_bb_kena_432d_field = 0;
    export let pasaran_win3d_bb_kena_432d_field = 0;
    export let pasaran_win3dd_bb_kena_432d_field = 0;
    export let pasaran_win2d_bb_kena_432d_field = 0;
    export let pasaran_win2dd_bb_kena_432d_field = 0;
    export let pasaran_win2dt_bb_kena_432d_field = 0;
    export let pasaran_win4d_bb_432d_field = 0;
    export let pasaran_win3d_bb_432d_field = 0;
    export let pasaran_win3dd_bb_432d_field = 0;
    export let pasaran_win2d_bb_432d_field = 0;
    export let pasaran_win2dd_bb_432d_field = 0;
    export let pasaran_win2dt_bb_432d_field = 0;
    let buttonLoading_class = "btn btn-primary";
    let msg_error = "";
    let tab_config_1 = "bg-sky-600 text-white"
    let tab_config_2 = ""
    let panel_conf_1 = true
    let panel_conf_2 = false
    let dispatch = createEventDispatcher();
    async function save432d() {
        let flag = false;
        msg_error = "";
        if (isNaN(pasaran_minbet_432d_field)) {
            flag = true;
            msg_error += "The Min Bet is required<br>";
        }
        if (pasaran_maxbet4d_432d_field == "") {
            flag = true;
            msg_error += "The Max Bet 4D is required<br>";
        }
        if (pasaran_maxbet3d_432d_field == "") {
            flag = true;
            msg_error += "The Max Bet 3D is required<br>";
        }
        if (pasaran_maxbet3dd_432d_field == "") {
            flag = true;
            msg_error += "The Max Bet 3DD is required<br>";
        }
        if (pasaran_maxbet2d_432d_field == "") {
            flag = true;
            msg_error += "The Max Bet 2D is required<br>";
        }
        if (pasaran_maxbet2dd_432d_field == "") {
            flag = true;
            msg_error += "The Max Bet 2DD is required<br>";
        }
        if (pasaran_maxbet2dt_432d_field == "") {
            flag = true;
            msg_error += "The Max Bet 2DT is required<br>";
        }
        if (pasaran_limitotal4d_432d_field == "") {
            flag = true;
            msg_error += "The Limit Total 4D is required<br>";
        }
        if (pasaran_limitotal3d_432d_field == "") {
            flag = true;
            msg_error += "The Limit Total 3D is required<br>";
        }
        if (pasaran_limitotal3dd_432d_field == "") {
            flag = true;
            msg_error += "The Limit Total 3DD is required<br>";
        }
        if (pasaran_limitotal2d_432d_field == "") {
            flag = true;
            msg_error += "The Limit Total 2D is required<br>";
        }
        if (pasaran_limitotal2dd_432d_field == "") {
            flag = true;
            msg_error += "The Limit Total 2DD is required<br>";
        }
        if (pasaran_limitotal2dt_432d_field == "") {
            flag = true;
            msg_error += "The Limit Total 2DT is required<br>";
        }
        if (pasaran_limitglobal4d_432d_field == "") {
            flag = true;
            msg_error += "The Limit Global 4D is required<br>";
        }
        if (pasaran_limitglobal3d_432d_field == "") {
            flag = true;
            msg_error += "The Limit Global 3D is required<br>";
        }
        if (pasaran_limitglobal3dd_432d_field == "") {
            flag = true;
            msg_error += "The Limit Global 3DD is required<br>";
        }
        if (pasaran_limitglobal2d_432d_field == "") {
            flag = true;
            msg_error += "The Limit Global 3D is required<br>";
        }
        if (pasaran_limitglobal2dd_432d_field == "") {
            flag = true;
            msg_error += "The Limit Global 2DD is required<br>";
        }
        if (pasaran_limitglobal2dt_432d_field == "") {
            flag = true;
            msg_error += "The Limit Global 2DT is required<br>";
        }
        if (pasaran_win4d_432d_field == "") {
            flag = true;
            msg_error += "The Win 4D is required<br>";
        }
        if (pasaran_win3d_432d_field == "") {
            flag = true;
            msg_error += "The Win 3D is required<br>";
        }
        if (pasaran_win3dd_432d_field == "") {
            flag = true;
            msg_error += "The Win 3DD is required<br>";
        }
        if (pasaran_win2d_432d_field == "") {
            flag = true;
            msg_error += "The Win 2D is required<br>";
        }
        if (pasaran_win2dd_432d_field == "") {
            flag = true;
            msg_error += "The Win 2DD is required<br>";
        }
        if (pasaran_win2dt_432d_field == "") {
            flag = true;
            msg_error += "The Win 2DT is required<br>";
        }
        if (pasaran_win4d_bb_kena_432d_field == "") {
            flag = true;
            msg_error += "The Win 4D BB KENA is required<br>";
        }
        if (pasaran_win3d_bb_kena_432d_field == "") {
            flag = true;
            msg_error += "The Win 3D BB KENA is required<br>";
        }
        if (pasaran_win3dd_bb_kena_432d_field == "") {
            flag = true;
            msg_error += "The Win 3DD BB KENA is required<br>";
        }
        if (pasaran_win2d_bb_kena_432d_field == "") {
            flag = true;
            msg_error += "The Win 2D BB KENA is required<br>";
        }
        if (pasaran_win2dd_bb_kena_432d_field == "") {
            flag = true;
            msg_error += "The Win 2DD BB KENA is required<br>";
        }
        if (pasaran_win2dt_bb_kena_432d_field == "") {
            flag = true;
            msg_error += "The Win 2DT BB KENA is required<br>";
        }
        if (pasaran_win4d_bb_432d_field == "") {
            flag = true;
            msg_error += "The Win 4D BB is required<br>";
        }
        if (pasaran_win3d_bb_432d_field == "") {
            flag = true;
            msg_error += "The Win 3D BB is required<br>";
        }
        if (pasaran_win3dd_bb_432d_field == "") {
            flag = true;
            msg_error += "The Win 3DD BB is required<br>";
        }
        if (pasaran_win2d_bb_432d_field == "") {
            flag = true;
            msg_error += "The Win 2D BB is required<br>";
        }
        if (pasaran_win2dd_bb_432d_field == "") {
            flag = true;
            msg_error += "The Win 2DD BB is required<br>";
        }
        if (pasaran_win2dt_bb_432d_field == "") {
            flag = true;
            msg_error += "The Win 2DT BB is required<br>";
        }
        if (pasaran_disc4d_432d_field == "") {
            flag = true;
            msg_error += "The Disc 4D is required<br>";
        }
        if (pasaran_disc3d_432d_field == "") {
            flag = true;
            msg_error += "The Disc 3D is required<br>";
        }
        if (pasaran_disc3dd_432d_field == "") {
            flag = true;
            msg_error += "The Disc 3DD is required<br>";
        }
        if (pasaran_disc2d_432d_field == "") {
            flag = true;
            msg_error += "The Disc 2D is required<br>";
        }
        if (pasaran_disc2dd_432d_field == "") {
            flag = true;
            msg_error += "The Disc 2DD is required<br>";
        }
        if (pasaran_disc2dt_432d_field == "") {
            flag = true;
            msg_error += "The Disc 2DT is required<br>";
        }
        if (flag == false) {
            buttonLoading_class = "btn loading"
            dispatch("handleLoadingRunning", "call");
            const res = await fetch(path_api+"api/savepasaranconf432", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    page: "PASARAN-SAVE",
                    idpasarantogel: pasaran_idpasarantogel_field,
                    pasaran_minbet_432d: parseInt(pasaran_minbet_432d_field),
                    pasaran_maxbet4d_432d: parseInt(pasaran_maxbet4d_432d_field),
                    pasaran_maxbet3d_432d: parseInt(pasaran_maxbet3d_432d_field),
                    pasaran_maxbet3dd_432d: parseInt(pasaran_maxbet3dd_432d_field),
                    pasaran_maxbet2d_432d: parseInt(pasaran_maxbet2d_432d_field),
                    pasaran_maxbet2dd_432d: parseInt(pasaran_maxbet2dd_432d_field),
                    pasaran_maxbet2dt_432d: parseInt(pasaran_maxbet2dt_432d_field),
                    pasaran_limitotal4d_432d: parseInt(pasaran_limitotal4d_432d_field),
                    pasaran_limitotal3d_432d: parseInt(pasaran_limitotal3d_432d_field),
                    pasaran_limitotal3dd_432d: parseInt(pasaran_limitotal3dd_432d_field),
                    pasaran_limitotal2d_432d: parseInt(pasaran_limitotal2d_432d_field),
                    pasaran_limitotal2dd_432d: parseInt(pasaran_limitotal2dd_432d_field),
                    pasaran_limitotal2dt_432d: parseInt(pasaran_limitotal2dt_432d_field),
                    pasaran_limitglobal4d_432d: parseInt(pasaran_limitglobal4d_432d_field),
                    pasaran_limitglobal3d_432d: parseInt(pasaran_limitglobal3d_432d_field),
                    pasaran_limitglobal3dd_432d: parseInt(pasaran_limitglobal3dd_432d_field),
                    pasaran_limitglobal2d_432d: parseInt(pasaran_limitglobal2d_432d_field),
                    pasaran_limitglobal2dd_432d: parseInt(pasaran_limitglobal2dd_432d_field),
                    pasaran_limitglobal2dt_432d: parseInt(pasaran_limitglobal2dt_432d_field),
                    pasaran_win4d_432d: parseInt(pasaran_win4d_432d_field),
                    pasaran_win3d_432d: parseInt(pasaran_win3d_432d_field),
                    pasaran_win3dd_432d: parseInt(pasaran_win3dd_432d_field),
                    pasaran_win2d_432d: parseInt(pasaran_win2d_432d_field),
                    pasaran_win2dd_432d: parseInt(pasaran_win2dd_432d_field),
                    pasaran_win2dt_432d: parseInt(pasaran_win2dt_432d_field),
                    pasaran_win4dnodisc_432d: parseInt(pasaran_win4dnodisc_432d_field),
                    pasaran_win3dnodisc_432d: parseInt(pasaran_win3dnodisc_432d_field),
                    pasaran_win3ddnodisc_432d: parseInt(pasaran_win3ddnodisc_432d_field),
                    pasaran_win2dnodisc_432d: parseInt(pasaran_win2dnodisc_432d_field),
                    pasaran_win2ddnodisc_432d: parseInt(pasaran_win2ddnodisc_432d_field),
                    pasaran_win2dtnodisc_432d: parseInt(pasaran_win2dtnodisc_432d_field),
                    pasaran_win4d_bb_kena_432d: parseInt(pasaran_win4d_bb_kena_432d_field),
                    pasaran_win3d_bb_kena_432d: parseInt(pasaran_win3d_bb_kena_432d_field),
                    pasaran_win3dd_bb_kena_432d: parseInt(pasaran_win3dd_bb_kena_432d_field),
                    pasaran_win2d_bb_kena_432d: parseInt(pasaran_win2d_bb_kena_432d_field),
                    pasaran_win2dd_bb_kena_432d: parseInt(pasaran_win2dd_bb_kena_432d_field),
                    pasaran_win2dt_bb_kena_432d: parseInt(pasaran_win2dt_bb_kena_432d_field),
                    pasaran_win4d_bb_432d: parseInt(pasaran_win4d_bb_432d_field),
                    pasaran_win3d_bb_432d: parseInt(pasaran_win3d_bb_432d_field),
                    pasaran_win3dd_bb_432d: parseInt(pasaran_win3dd_bb_432d_field),
                    pasaran_win2d_bb_432d: parseInt(pasaran_win2d_bb_432d_field),
                    pasaran_win2dd_bb_432d: parseInt(pasaran_win2dd_bb_432d_field),
                    pasaran_win2dt_bb_432d: parseInt(pasaran_win2dt_bb_432d_field),
                    pasaran_disc4d_432d: parseFloat(pasaran_disc4d_432d_field / 100),
                    pasaran_disc3d_432d: parseFloat(pasaran_disc3d_432d_field / 100),
                    pasaran_disc3dd_432d: parseFloat(pasaran_disc3dd_432d_field / 100),
                    pasaran_disc2d_432d: parseFloat(pasaran_disc2d_432d_field / 100),
                    pasaran_disc2dd_432d: parseFloat(pasaran_disc2dd_432d_field / 100),
                    pasaran_disc2dt_432d: parseFloat( pasaran_disc2dt_432d_field / 100),
                }),
            });
            const json = await res.json();
            if(!res.ok){
                let temp_msg = "System Mengalami Trouble"
                dispatch("handleLoadingRunningFinish", {
                        temp_msg
                });
            }else{
                let temp_msg = json.message
                dispatch("handleLoadingRunningFinish", {
                        temp_msg
                });
            }
            buttonLoading_class = "btn btn-primary"
        } else {
            if(msg_error != ""){
                let temp_msg = msg_error
                dispatch("handleCallNotif", {
                        temp_msg
                });
            }
        }
    }
    const ChangeTabMenu = (e) => {
        switch(e){
            case "config_1":
                tab_config_1 = "bg-sky-600 text-white"
                tab_config_2 = ""
                panel_conf_1 = true
                panel_conf_2 = false
                break;
            case "config_2":
                tab_config_1 = ""
                tab_config_2 = "bg-sky-600 text-white"
                panel_conf_1 = false
                panel_conf_2 = true
                break;
        }
    }
</script>

<ul class="flex justify-center items-center gap-2">
    <li on:click={() => {
            ChangeTabMenu("config_1");
        }}
        class="items-center {tab_config_1}  px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">Configure 1</li>
    <li on:click={() => {
            ChangeTabMenu("config_2");
        }}
        class="items-center {tab_config_2} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">Configure 2</li>
</ul>
{#if panel_conf_1}
    <div class="grid grid-cols-3 gap-1 mt-5 mb-5">
        <Input_custom
                input_enabled={true}
                input_tipe="number"
                bind:value={pasaran_minbet_432d_field}
                input_id="pasaran_minbet_432d_field"
                input_placeholder="Minimal Bet"/>
        <div class="col-span-2"> </div>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_maxbet4d_432d_field}
            input_id="pasaran_maxbet4d_432d_field"
            input_placeholder="Max Bet 4D"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_limitotal4d_432d_field}
            input_id="pasaran_limitotal4d_432d_field"
            input_placeholder="Limit Total 4D"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_limitglobal4d_432d_field}
            input_id="pasaran_limitotal4d_432d_field"
            input_placeholder="Limit Global 4D"/>
        
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_maxbet3d_432d_field}
            input_id="pasaran_maxbet3d_432d_field"
            input_placeholder="Max Bet 3D"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_limitotal3d_432d_field}
            input_id="pasaran_limitotal3d_432d_field"
            input_placeholder="Limit Total 3D"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_limitglobal3d_432d_field}
            input_id="pasaran_limitglobal3d_432d_field"
            input_placeholder="Limit Global 3D"/>
        
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_maxbet3dd_432d_field}
            input_id="pasaran_maxbet3dd_432d_field"
            input_placeholder="Max Bet 3DD"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_limitotal3dd_432d_field}
            input_id="pasaran_limitotal3dd_432d_field"
            input_placeholder="Limit Total 3DD"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_limitglobal3dd_432d_field}
            input_id="pasaran_limitglobal3dd_432d_field"
            input_placeholder="Limit Global 3DD"/>
        
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_maxbet2d_432d_field}
            input_id="pasaran_maxbet2d_432d_field"
            input_placeholder="Max Bet 2D"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_limitotal2d_432d_field}
            input_id="pasaran_limitotal2d_432d_field"
            input_placeholder="Limit Total 2D"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_limitglobal2d_432d_field}
            input_id="pasaran_limitglobal2d_432d_field"
            input_placeholder="Limit Global 2D"/>
        
        
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_maxbet2dd_432d_field}
            input_id="pasaran_maxbet2dd_432d_field"
            input_placeholder="Max Bet 2DD"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_limitotal2dd_432d_field}
            input_id="pasaran_limitotal2dd_432d_field"
            input_placeholder="Limit Total 2DD"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_limitglobal2dd_432d_field}
            input_id="pasaran_limitotal2dd_432d_field"
            input_placeholder="Limit Global 2DD"/>
        
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_maxbet2dt_432d_field}
            input_id="pasaran_maxbet2dt_432d_field"
            input_placeholder="Max Bet 2DT"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_limitotal2dt_432d_field}
            input_id="pasaran_limitotal2dt_432d_field"
            input_placeholder="Limit Total 2DT"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_limitglobal2dt_432d_field}
            input_id="pasaran_limitglobal2dt_432d_field"
            input_placeholder="Limit Global 2DT"/>
        
        
    </div>
{/if}
{#if panel_conf_2}
    <div class="grid grid-cols-5 gap-1 mt-5 mb-5">
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win4d_432d_field}
            input_id="pasaran_win4d_432d_field"
            input_placeholder="WIN 4D"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win4dnodisc_432d_field}
            input_id="pasaran_win4dnodisc_432d_field"
            input_placeholder="WIN 4D FULL(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win4d_bb_kena_432d_field}
            input_id="pasaran_win4d_bb_kena_432d_field"
            input_placeholder="WIN 4D BB - KENA(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win4d_bb_432d_field}
            input_id="pasaran_win4d_bb_kena_432d_field"
            input_placeholder="WIN 4D BB(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="float"
            bind:value={pasaran_disc4d_432d_field}
            input_precision=2
            input_id="pasaran_disc4d_432d_field"
            input_placeholder="DISC 4D(%)"/>
        
        
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win3d_432d_field}
            input_id="pasaran_win3d_432d_field"
            input_placeholder="WIN 3D"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win3dnodisc_432d_field}
            input_id="pasaran_win3d_432d_field"
            input_placeholder="WIN 3D FULL(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win3d_bb_kena_432d_field}
            input_id="pasaran_win3d_bb_kena_432d_field"
            input_placeholder="WIN 3D BB - KENA(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win3d_bb_432d_field}
            input_id="pasaran_win3d_bb_432d_field"
            input_placeholder="WIN 3D BB(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="float"
            bind:value={pasaran_disc3d_432d_field}
            input_precision=2
            input_id="pasaran_disc3d_432d_field"
            input_placeholder="DISC 3D(%)"/>
        
        
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win3dd_432d_field}
            input_id="pasaran_win3dd_432d_field"
            input_placeholder="WIN 3DD"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win3ddnodisc_432d_field}
            input_id="pasaran_win3ddnodisc_432d_field"
            input_placeholder="WIN 3DD FULL(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win3dd_bb_kena_432d_field}
            input_id="pasaran_win3dd_bb_kena_432d_field"
            input_placeholder="WIN 3DD BB - KENA(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win3dd_bb_432d_field}
            input_id="pasaran_win3dd_bb_432d_field"
            input_placeholder="WIN 3DD BB(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="float"
            bind:value={pasaran_disc3dd_432d_field}
            input_precision=2
            input_id="pasaran_disc3dd_432d_field"
            input_placeholder="DISC 3DD(%)"/>
        
        
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win2d_432d_field}
            input_id="pasaran_win2d_432d_field"
            input_placeholder="WIN 2D"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win2dnodisc_432d_field}
            input_id="pasaran_win2dnodisc_432d_field"
            input_placeholder="WIN 2D FULL(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win2d_bb_kena_432d_field}
            input_id="pasaran_win2d_bb_kena_432d_field"
            input_placeholder="WIN 2D BB - KENA(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win2d_bb_432d_field}
            input_id="pasaran_win2d_bb_432d_field"
            input_placeholder="WIN 2D BB(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="float"
            bind:value={pasaran_disc2d_432d_field}
            input_precision=2
            input_id="pasaran_disc2d_432d_field"
            input_placeholder="DISC 2D(%)"/>
        
        
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win2dd_432d_field}
            input_id="pasaran_win2dd_432d_field"
            input_placeholder="WIN 2DD"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win2ddnodisc_432d_field}
            input_id="pasaran_win2dd_432d_field"
            input_placeholder="WIN 2DD FULL(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win2dd_bb_kena_432d_field}
            input_id="pasaran_win2dd_bb_kena_432d_field"
            input_placeholder="WIN 2DD BB - KENA(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win2dd_bb_432d_field}
            input_id="pasaran_win2dd_bb_432d_field"
            input_placeholder="WIN 2DD BB(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="float"
            bind:value={pasaran_disc2dd_432d_field}
            input_precision=2
            input_id="pasaran_disc2dd_432d_field"
            input_placeholder="DISC 2DD(%)"/>
        
        
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win2dt_432d_field}
            input_id="pasaran_win2dt_432d_field"
            input_placeholder="WIN 2DT"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win2dtnodisc_432d_field}
            input_id="pasaran_win2dtnodisc_432d_field"
            input_placeholder="WIN 2DT FULL(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win2dt_bb_kena_432d_field}
            input_id="pasaran_win2dt_bb_kena_432d_field"
            input_placeholder="WIN 2DT BB - KENA(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_win2dt_bb_432d_field}
            input_id="pasaran_win2dt_bb_432d_field"
            input_placeholder="WIN 2DT BB(x)"/>
        <Input_custom
            input_enabled={true}
            input_tipe="float"
            bind:value={pasaran_disc2dt_432d_field}
            input_precision=2
            input_id="pasaran_disc2dt_432d_field"
            input_placeholder="DISC 2DT(%)"/>
    </div>
{/if}
{#if pasaran_tipe != "WAJIB"}
<button on:click={() => {
    save432d();
}} class="{buttonLoading_class} btn-block ">Submit</button>
{/if}