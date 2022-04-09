<script>
    import { createEventDispatcher } from "svelte";
    import Input_custom from '../../components/Input.svelte'
    export let pasaran_tipe = "";
    export let path_api = "";
    export let token = "";
    export let idcomppasaran = "";
    export let pasaran_idpasarantogel_field = "";
    export let pasaran_minbet_cbebas_field = 0;
    export let pasaran_maxbet_cbebas_field = 0;
    export let pasaran_limitotal_cbebas_field = 0;
    export let pasaran_limitglobal_cbebas_field = 0;
    export let pasaran_win_cbebas_field = 0;
    export let pasaran_disc_cbebas_field = 0;
    let buttonLoading_class = "btn btn-primary";
    let msg_error = "";
    let dispatch = createEventDispatcher();
    async function save432d() {
        let flag = false;
        msg_error = "";
        if (pasaran_minbet_cbebas_field == "") {
            flag = true;
            msg_error += "The Min Bet is required<br>";
        }
        if (pasaran_maxbet_cbebas_field == "") {
            flag = true;
            msg_error += "The Max Bet is required<br>";
        }
        if (pasaran_limitotal_cbebas_field == "") {
            flag = true;
            msg_error += "The Limit Total is required<br>";
        }
        if (pasaran_limitglobal_cbebas_field == "") {
            flag = true;
            msg_error += "The Limit Global is required<br>";
        }
        if (pasaran_win_cbebas_field == "") {
            flag = true;
            msg_error += "The Win is required<br>";
        }
        if (pasaran_disc_cbebas_field == "") {
            flag = true;
            msg_error += "The Diskon is required<br>";
        }
        if (flag == false) {
            buttonLoading_class = "btn loading"
            dispatch("handleLoadingRunning", "call");
            const res = await fetch(path_api+"api/savepasaranconfcolokbebas", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    page: "PASARAN-SAVE",
                    idpasarantogel: pasaran_idpasarantogel_field,
                    pasaran_minbet_cbebas: parseInt(pasaran_minbet_cbebas_field),
                    pasaran_maxbet_cbebas: parseInt(pasaran_maxbet_cbebas_field),
                    pasaran_limitotal_cbebas: parseInt(pasaran_limitotal_cbebas_field),
                    pasaran_limitglobal_cbebas: parseInt(pasaran_limitglobal_cbebas_field),
                    pasaran_win_cbebas: parseFloat(pasaran_win_cbebas_field),
                    pasaran_disc_cbebas: parseFloat(pasaran_disc_cbebas_field / 100),
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
    
</script>

<div class="grid grid-cols-3 gap-1 mt-2 mb-5">
    <Input_custom
            input_enabled={true}
            input_tipe="number"
            bind:value={pasaran_minbet_cbebas_field}
            input_id="pasaran_minbet_cbebas_field"
            input_placeholder="Minimal Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_limitotal_cbebas_field}
        input_id="pasaran_limitotal_cbebas_field"
        input_placeholder="Limit Total"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_win_cbebas_field}
        input_precision=2
        input_id="pasaran_win_cbebas_field"
        input_placeholder="WIN(x)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_maxbet_cbebas_field}
        input_id="pasaran_maxbet_cbebas_field"
        input_placeholder="Max Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_limitglobal_cbebas_field}
        input_id="pasaran_limitglobal_cbebas_field"
        input_placeholder="Limit Global"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_disc_cbebas_field}
        input_precision=2
        input_id="pasaran_disc_cbebas_field"
        input_placeholder="DISC(%)"/>
</div>
{#if pasaran_tipe != "WAJIB"}
<button on:click={() => {
    save432d();
}} class="{buttonLoading_class} btn-block ">Submit</button>
{/if}