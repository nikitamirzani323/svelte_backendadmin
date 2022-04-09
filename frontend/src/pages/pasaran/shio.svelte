<script>
    import { createEventDispatcher } from "svelte";
    import Input_custom from '../../components/Input.svelte'
    export let pasaran_tipe = "";
    export let path_api = "";
    export let token = "";
    export let idcomppasaran = "";
    export let pasaran_idpasarantogel_field = "";
    export let pasaran_minbet_shio_field = 0;
    export let pasaran_maxbet_shio_field = 0;
    export let pasaran_win_shio_field = 0;
    export let pasaran_disc_shio_field = 0;
    export let pasaran_shioyear_shio_field = "";
    export let pasaran_limitglobal_shio_field = 0;
    export let pasaran_limittotal_shio_field = 0;
    let buttonLoading_class = "btn btn-primary";
    let msg_error = "";
    let dispatch = createEventDispatcher();
    async function save432d() {
        let flag = false;
        msg_error = "";
        if (pasaran_minbet_shio_field == "") {
            flag = true;
            msg_error += "The Min Bet is required<br>";
        }
        if (pasaran_maxbet_shio_field == "") {
            flag = true;
            msg_error += "The Max Bet is required<br>";
        }
        if (pasaran_limittotal_shio_field == "") {
            flag = true;
            msg_error += "The Limit Total is required<br>";
        }
        if (pasaran_limitglobal_shio_field == "") {
            flag = true;
            msg_error += "The Limit Global is required<br>";
        }
        if (pasaran_win_shio_field == "") {
            flag = true;
            msg_error += "The Win is required<br>";
        }
        if (pasaran_disc_shio_field == "") {
            flag = true;
            msg_error += "The Diskon is required<br>";
        }
        if (flag == false) {
            buttonLoading_class = "btn loading"
            dispatch("handleLoadingRunning", "call");
            const res = await fetch(path_api+"api/savepasaranconfshio", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    page: "PASARAN-SAVE",
                    idpasarantogel: pasaran_idpasarantogel_field,
                    pasaran_minbet_shio: parseInt(pasaran_minbet_shio_field),
                    pasaran_maxbet_shio: parseInt(pasaran_maxbet_shio_field),
                    pasaran_limittotal_shio: parseInt(pasaran_limittotal_shio_field),
                    pasaran_limitglobal_shio: parseInt(pasaran_limitglobal_shio_field),
                    pasaran_shioyear_shio: "kerbau",
                    pasaran_disc_shio: parseFloat((pasaran_disc_shio_field / 100).toPrecision(3)),
                    pasaran_win_shio: parseFloat(pasaran_win_shio_field),
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

<div class="grid grid-cols-4 gap-1 mt-2 mb-5">
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_minbet_shio_field}
        input_id="pasaran_minbet_shio_field"
        input_placeholder="Minimal Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_limittotal_shio_field}
        input_id="pasaran_limittotal_shio_field"
        input_placeholder="Limit Total"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_win_shio_field}
        input_id="pasaran_win_shio_field"
        input_placeholder="WIN(x)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_disc_shio_field}
        input_id="pasaran_disc_shio_field"
        input_placeholder="DISC(%)"/>
    
    
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_maxbet_shio_field}
        input_id="pasaran_maxbet_shio_field"
        input_placeholder="Max Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_limitglobal_shio_field}
        input_id="pasaran_limitglobal_shio_field"
        input_placeholder="Limit Global"/>
    
</div>

{#if pasaran_tipe != "WAJIB"}
<button on:click={() => {
    save432d();
}} class="{buttonLoading_class} btn-block">Submit</button>
{/if}