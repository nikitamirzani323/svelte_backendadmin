<script>
    import { createEventDispatcher } from "svelte";
    import Input_custom from '../../components/Input.svelte'
    export let pasaran_tipe = "";
    export let path_api = "";
    export let token = "";
    export let idcomppasaran = "";
    export let pasaran_idpasarantogel_field = "";
    export let pasaran_minbet_dasar_field = 0;
    export let pasaran_maxbet_dasar_field = 0;
    export let pasaran_keibesar_dasar_field = 0;
    export let pasaran_keikecil_dasar_field = 0;
    export let pasaran_keigenap_dasar_field = 0;
    export let pasaran_keiganjil_dasar_field = 0;
    export let pasaran_discbesar_dasar_field = 0;
    export let pasaran_disckecil_dasar_field = 0;
    export let pasaran_discgenap_dasar_field = 0;
    export let pasaran_discganjil_dasar_field = 0;
    export let pasaran_limitglobal_dasar_field = 0;
    export let pasaran_limittotal_dasar_field = 0;
    let buttonLoading_class = "btn btn-primary";
    let msg_error = "";
    let dispatch = createEventDispatcher();
    async function save432d() {
        let flag = false;
        msg_error = "";
        if (pasaran_minbet_dasar_field == "") {
            flag = true;
            msg_error += "The Min Bet is required<br>";
        }
        if (pasaran_maxbet_dasar_field == "") {
            flag = true;
            msg_error += "The Max Bet is required<br>";
        }
        if (pasaran_limittotal_dasar_field == "") {
            flag = true;
            msg_error += "The Limit Total is required<br>";
        }
        if (pasaran_limitglobal_dasar_field == "") {
            flag = true;
            msg_error += "The Limit Global is required<br>";
        }
        if (pasaran_keibesar_dasar_field == "") {
            flag = true;
            msg_error += "The Kei Besar is required<br>";
        }
        if (pasaran_keikecil_dasar_field == "") {
            flag = true;
            msg_error += "The Kei Kecil is required<br>";
        }
        if (pasaran_keigenap_dasar_field == "") {
            flag = true;
            msg_error += "The Kei Genap is required<br>";
        }
        if (pasaran_keiganjil_dasar_field == "") {
            flag = true;
            msg_error += "The Kei Ganjil is required<br>";
        }
        if (pasaran_discbesar_dasar_field == "") {
            flag = true;
            msg_error += "The Diskon Besar is required<br>";
        }
        if (pasaran_disckecil_dasar_field == "") {
            flag = true;
            msg_error += "The Diskon Kecil is required<br>";
        }
        if (pasaran_discgenap_dasar_field == "") {
            flag = true;
            msg_error += "The Diskon Genap is required<br>";
        }
        if (pasaran_discganjil_dasar_field == "") {
            flag = true;
            msg_error += "The Diskon Ganjil is required<br>";
        }
        if (flag == false) {
            buttonLoading_class = "btn loading"
            dispatch("handleLoadingRunning", "call");
            const res = await fetch(path_api+"api/savepasaranconfdasar", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    page: "PASARAN-SAVE",
                    idpasarantogel: pasaran_idpasarantogel_field,
                    pasaran_minbet_dasar: parseInt(pasaran_minbet_dasar_field),
                    pasaran_maxbet_dasar: parseInt(pasaran_maxbet_dasar_field),
                    pasaran_limittotal_dasar: parseInt(pasaran_limittotal_dasar_field),
                    pasaran_limitglobal_dasar: parseInt(pasaran_limitglobal_dasar_field),
                    pasaran_keibesar_dasar: parseFloat((pasaran_keibesar_dasar_field / 100).toPrecision(3)),
                    pasaran_keikecil_dasar: parseFloat((pasaran_keikecil_dasar_field / 100).toPrecision(3)),
                    pasaran_keigenap_dasar: parseFloat((pasaran_keigenap_dasar_field / 100).toPrecision(3)),
                    pasaran_keiganjil_dasar: parseFloat((pasaran_keiganjil_dasar_field / 100).toPrecision(3)),
                    pasaran_discbesar_dasar: parseFloat((pasaran_discbesar_dasar_field / 100).toPrecision(3)),
                    pasaran_disckecil_dasar: parseFloat((pasaran_disckecil_dasar_field / 100).toPrecision(3)),
                    pasaran_discgenap_dasar: parseFloat((pasaran_discgenap_dasar_field / 100).toPrecision(3)),
                    pasaran_discganjil_dasar: parseFloat((pasaran_discganjil_dasar_field / 100).toPrecision(3)),
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
        bind:value={pasaran_minbet_dasar_field}
        input_id="pasaran_minbet_dasar_field"
        input_placeholder="Minimal Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_limittotal_dasar_field}
        input_id="pasaran_limittotal_dasar_field"
        input_placeholder="Limit Total"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_keibesar_dasar_field}
        input_id="pasaran_keibesar_dasar_field"
        input_placeholder="KEI BESAR(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_discbesar_dasar_field}
        input_id="pasaran_discbesar_dasar_field"
        input_placeholder="DISC BESAR(%)"/>
    
    
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_maxbet_dasar_field}
        input_id="pasaran_maxbet_dasar_field"
        input_placeholder="Max Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_limitglobal_dasar_field}
        input_id="pasaran_limitglobal_dasar_field"
        input_placeholder="Limit Global"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_keikecil_dasar_field}
        input_id="pasaran_keikecil_dasar_field"
        input_placeholder="KEI KECIL(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_disckecil_dasar_field}
        input_id="pasaran_disckecil_dasar_field"
        input_placeholder="DISC KECIL(%)"/>
    
    <div class="col-span-2"></div>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_keigenap_dasar_field}
        input_id="pasaran_keigenap_dasar_field"
        input_placeholder="KEI GENAP(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_discgenap_dasar_field}
        input_id="pasaran_discgenap_dasar_field"
        input_placeholder="DISC GENAP(%)"/>
    
    <div class="col-span-2"></div>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_keiganjil_dasar_field}
        input_id="pasaran_keiganjil_dasar_field"
        input_placeholder="KEI GANJIL(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_discganjil_dasar_field}
        input_id="pasaran_discganjil_dasar_field"
        input_placeholder="DISC GANJIL(%)"/>
</div>
{#if pasaran_tipe != "WAJIB"}
<button on:click={() => {
    save432d();
}} class="{buttonLoading_class} btn-block">Submit</button>
{/if}