<script>
    import { createEventDispatcher } from "svelte";
    import Input_custom from '../../components/Input.svelte'
    export let pasaran_tipe = "";
    export let path_api = "";
    export let token = "";
    export let idcomppasaran = "";
    export let pasaran_idpasarantogel_field = "";
    export let pasaran_minbet_5050special_field = 0;
    export let pasaran_maxbet_5050special_field = 0;
    export let pasaran_keiasganjil_5050special_field = 0;
    export let pasaran_keiasgenap_5050special_field = 0;
    export let pasaran_keiasbesar_5050special_field = 0;
    export let pasaran_keiaskecil_5050special_field = 0;
    export let pasaran_keikopganjil_5050special_field = 0;
    export let pasaran_keikopgenap_5050special_field = 0;
    export let pasaran_keikopbesar_5050special_field = 0;
    export let pasaran_keikopkecil_5050special_field = 0;
    export let pasaran_keikepalaganjil_5050special_field = 0;
    export let pasaran_keikepalagenap_5050special_field = 0;
    export let pasaran_keikepalabesar_5050special_field = 0;
    export let pasaran_keikepalakecil_5050special_field = 0;
    export let pasaran_keiekorganjil_5050special_field = 0;
    export let pasaran_keiekorgenap_5050special_field = 0;
    export let pasaran_keiekorbesar_5050special_field = 0;
    export let pasaran_keiekorkecil_5050special_field = 0;
    export let pasaran_discasganjil_5050special_field = 0;
    export let pasaran_discasgenap_5050special_field = 0;
    export let pasaran_discasbesar_5050special_field = 0;
    export let pasaran_discaskecil_5050special_field = 0;
    export let pasaran_disckopganjil_5050special_field = 0;
    export let pasaran_disckopgenap_5050special_field = 0;
    export let pasaran_disckopbesar_5050special_field = 0;
    export let pasaran_disckopkecil_5050special_field = 0;
    export let pasaran_disckepalaganjil_5050special_field = 0;
    export let pasaran_disckepalagenap_5050special_field = 0;
    export let pasaran_disckepalabesar_5050special_field = 0;
    export let pasaran_disckepalakecil_5050special_field = 0;
    export let pasaran_discekorganjil_5050special_field = 0;
    export let pasaran_discekorgenap_5050special_field = 0;
    export let pasaran_discekorbesar_5050special_field = 0;
    export let pasaran_discekorkecil_5050special_field = 0;
    export let pasaran_limitglobal_5050special_field = 0;
    export let pasaran_limittotal_5050special_field = 0;
    let buttonLoading_class = "btn btn-primary";
    let msg_error = "";
    let dispatch = createEventDispatcher();
    async function save432d() {
        let flag = false;
        msg_error = "";
        if (pasaran_minbet_5050special_field == "") {
            flag = true;
            msg_error += "The Min Bet is required<br>";
        }
        if (pasaran_maxbet_5050special_field == "") {
            flag = true;
            msg_error += "The Max Bet is required<br>";
        }
        if (pasaran_limittotal_5050special_field == "") {
            flag = true;
            msg_error += "The Limit Total is required<br>";
        }
        if (pasaran_limitglobal_5050special_field == "") {
            flag = true;
            msg_error += "The Limit Global is required<br>";
        }
        if (pasaran_keiasganjil_5050special_field == "") {
            flag = true;
            msg_error += "The Kei AS Ganjil is required<br>";
        }
        if (pasaran_keiasgenap_5050special_field == "") {
            flag = true;
            msg_error += "The Kei As Genap is required<br>";
        }
        if (pasaran_keiasbesar_5050special_field == "") {
            flag = true;
            msg_error += "The Kei As Besar is required<br>";
        }
        if (pasaran_keiaskecil_5050special_field == "") {
            flag = true;
            msg_error += "The Kei As Kecil is required<br>";
        }
        if (pasaran_keikopganjil_5050special_field == "") {
            flag = true;
            msg_error += "The Kei Kop Ganjil is required<br>";
        }
        if (pasaran_keikopgenap_5050special_field == "") {
            flag = true;
            msg_error += "The Kei Kop Genap is required<br>";
        }
        if (pasaran_keikopbesar_5050special_field == "") {
            flag = true;
            msg_error += "The Kei Kop Besar is required<br>";
        }
        if (pasaran_keikopkecil_5050special_field == "") {
            flag = true;
            msg_error += "The Kei Kop Kecil is required<br>";
        }
        if (pasaran_keikepalaganjil_5050special_field == "") {
            flag = true;
            msg_error += "The Kei Kepala Ganjil is required<br>";
        }
        if (pasaran_keikepalagenap_5050special_field == "") {
            flag = true;
            msg_error += "The Kei Kepala Genap is required<br>";
        }
        if (pasaran_keikepalabesar_5050special_field == "") {
            flag = true;
            msg_error += "The Kei Kepala Besar is required<br>";
        }
        if (pasaran_keikepalakecil_5050special_field == "") {
            flag = true;
            msg_error += "The Kei Kepala Kecil is required<br>";
        }
        if (pasaran_keiekorganjil_5050special_field == "") {
            flag = true;
            msg_error += "The Kei Ekor Ganjil is required<br>";
        }
        if (pasaran_keiekorgenap_5050special_field == "") {
            flag = true;
            msg_error += "The Kei Ekor Genap is required<br>";
        }
        if (pasaran_keiekorbesar_5050special_field == "") {
            flag = true;
            msg_error += "The Kei Ekor Besar is required<br>";
        }
        if (pasaran_keiekorkecil_5050special_field == "") {
            flag = true;
            msg_error += "The Kei Ekor Kecil is required<br>";
        }
        if (pasaran_discasganjil_5050special_field == "") {
            flag = true;
            msg_error += "The Disc As Ganjil is required<br>";
        }
        if (pasaran_discasgenap_5050special_field == "") {
            flag = true;
            msg_error += "The Disc As Genap is required<br>";
        }
        if (pasaran_discasbesar_5050special_field == "") {
            flag = true;
            msg_error += "The Disc As Besar is required<br>";
        }
        if (pasaran_discaskecil_5050special_field == "") {
            flag = true;
            msg_error += "The Disc As Kecil is required<br>";
        }
        if (pasaran_disckopganjil_5050special_field == "") {
            flag = true;
            msg_error += "The Disc Kop Ganjil is required<br>";
        }
        if (pasaran_disckopgenap_5050special_field == "") {
            flag = true;
            msg_error += "The Disc Kop Genap is required<br>";
        }
        if (pasaran_disckopbesar_5050special_field == "") {
            flag = true;
            msg_error += "The Disc Kop Besar is required<br>";
        }
        if (pasaran_disckopkecil_5050special_field == "") {
            flag = true;
            msg_error += "The Disc Kop Kecil is required<br>";
        }
        if (pasaran_disckepalaganjil_5050special_field == "") {
            flag = true;
            msg_error += "The Disc Kepala Ganjil is required<br>";
        }
        if (pasaran_disckepalagenap_5050special_field == "") {
            flag = true;
            msg_error += "The Disc Kepala Genap is required<br>";
        }
        if (pasaran_disckepalabesar_5050special_field == "") {
            flag = true;
            msg_error += "The Disc Kepala Besar is required<br>";
        }
        if (pasaran_disckepalakecil_5050special_field == "") {
            flag = true;
            msg_error += "The Disc Kepala Kecil is required<br>";
        }
        if (pasaran_discekorganjil_5050special_field == "") {
            flag = true;
            msg_error += "The Disc Ekor Ganjil is required<br>";
        }
        if (pasaran_discekorgenap_5050special_field == "") {
            flag = true;
            msg_error += "The Disc Ekor Genap is required<br>";
        }
        if (pasaran_discekorbesar_5050special_field == "") {
            flag = true;
            msg_error += "The Disc Ekor Besar is required<br>";
        }
        if (pasaran_discekorkecil_5050special_field == "") {
            flag = true;
            msg_error += "The Disc Ekor Kecil is required<br>";
        }
        if (flag == false) {
            buttonLoading_class = "btn loading"
            dispatch("handleLoadingRunning", "call");
            const res = await fetch(path_api+"api/savepasaranconf5050special", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    page: "PASARAN-SAVE",
                    idpasarantogel: pasaran_idpasarantogel_field,
                    pasaran_minbet_5050special: parseInt(pasaran_minbet_5050special_field),
                    pasaran_maxbet_5050special: parseInt(pasaran_maxbet_5050special_field),
                    pasaran_limitglobal_5050special: parseInt(pasaran_limitglobal_5050special_field),
                    pasaran_limittotal_5050special: parseInt(pasaran_limittotal_5050special_field),
                    pasaran_keiasganjil_5050special: parseFloat(
                        pasaran_keiasganjil_5050special_field / 100
                    ),
                    pasaran_keiasgenap_5050special: parseFloat(
                        pasaran_keiasgenap_5050special_field / 100
                    ),
                    pasaran_keiasbesar_5050special: parseFloat(
                        pasaran_keiasbesar_5050special_field / 100
                    ),
                    pasaran_keiaskecil_5050special: parseFloat(
                        pasaran_keiaskecil_5050special_field / 100
                    ),
                    pasaran_keikopganjil_5050special: parseFloat(
                        pasaran_keikopganjil_5050special_field / 100
                    ),
                    pasaran_keikopgenap_5050special: parseFloat(
                        pasaran_keikopgenap_5050special_field / 100
                    ),
                    pasaran_keikopbesar_5050special: parseFloat(
                        pasaran_keikopbesar_5050special_field / 100
                    ),
                    pasaran_keikopkecil_5050special: parseFloat(
                        pasaran_keikopkecil_5050special_field / 100
                    ),
                    pasaran_keikepalaganjil_5050special: parseFloat(
                        pasaran_keikepalaganjil_5050special_field / 100
                    ),
                    pasaran_keikepalagenap_5050special: parseFloat(
                        pasaran_keikepalagenap_5050special_field / 100
                    ),
                    pasaran_keikepalabesar_5050special: parseFloat(
                        pasaran_keikepalabesar_5050special_field / 100
                    ),
                    pasaran_keikepalakecil_5050special: parseFloat(
                        pasaran_keikepalakecil_5050special_field / 100
                    ),
                    pasaran_keiekorganjil_5050special: parseFloat(
                        pasaran_keiekorganjil_5050special_field / 100
                    ),
                    pasaran_keiekorgenap_5050special: parseFloat(
                        pasaran_keiekorgenap_5050special_field / 100
                    ),
                    pasaran_keiekorbesar_5050special: parseFloat(
                        pasaran_keiekorbesar_5050special_field / 100
                    ),
                    pasaran_keiekorkecil_5050special: parseFloat(
                        pasaran_keiekorkecil_5050special_field / 100
                    ),
                    pasaran_discasganjil_5050special: parseFloat(
                        pasaran_discasganjil_5050special_field / 100
                    ),
                    pasaran_discasgenap_5050special: parseFloat(
                        pasaran_discasgenap_5050special_field / 100
                    ),
                    pasaran_discasbesar_5050special: parseFloat(
                        pasaran_discasbesar_5050special_field / 100
                    ),
                    pasaran_discaskecil_5050special: parseFloat(
                        pasaran_discaskecil_5050special_field / 100
                    ),
                    pasaran_disckopganjil_5050special: parseFloat(
                        pasaran_disckopganjil_5050special_field / 100
                    ),
                    pasaran_disckopgenap_5050special: parseFloat(
                        pasaran_disckopgenap_5050special_field / 100
                    ),
                    pasaran_disckopbesar_5050special: parseFloat(
                        pasaran_disckopbesar_5050special_field / 100
                    ),
                    pasaran_disckopkecil_5050special: parseFloat(
                        pasaran_disckopkecil_5050special_field / 100
                    ),
                    pasaran_disckepalaganjil_5050special: parseFloat(
                        pasaran_disckepalaganjil_5050special_field / 100
                    ),
                    pasaran_disckepalagenap_5050special: parseFloat(
                        pasaran_disckepalagenap_5050special_field / 100
                    ),
                    pasaran_disckepalabesar_5050special: parseFloat(
                        pasaran_disckepalabesar_5050special_field / 100
                    ),
                    pasaran_disckepalakecil_5050special: parseFloat(
                        pasaran_disckepalakecil_5050special_field / 100
                    ),
                    pasaran_discekorganjil_5050special: parseFloat(
                        pasaran_discekorganjil_5050special_field / 100
                    ),
                    pasaran_discekorgenap_5050special: parseFloat(
                        pasaran_discekorgenap_5050special_field / 100
                    ),
                    pasaran_discekorbesar_5050special: parseFloat(
                        pasaran_discekorbesar_5050special_field / 100
                    ),
                    pasaran_discekorkecil_5050special: parseFloat(
                        pasaran_discekorkecil_5050special_field / 100
                    ),
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
        bind:value={pasaran_minbet_5050special_field}
        input_id="pasaran_minbet_5050special_field"
        input_placeholder="Minimal Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_limittotal_5050special_field}
        input_id="pasaran_limittotal_5050special_field"
        input_placeholder="Limit Total"/>
    
    <div class="col-span-2"></div>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_maxbet_5050special_field}
        input_id="pasaran_maxbet_5050special_field"
        input_placeholder="Max Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        bind:value={pasaran_limitglobal_5050special_field}
        input_id="pasaran_limitglobal_5050special_field"
        input_placeholder="Limit Global"/>
    
    <div class="col-span-2"></div>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keiasganjil_5050special_field}
        input_precision=2
        input_id="pasaran_keiasganjil_5050special_field"
        input_placeholder="KEI AS GANJIL(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_discasganjil_5050special_field}
        input_precision=2
        input_id="pasaran_discasganjil_5050special_field"
        input_placeholder="DISC AS GANJIL(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keiasgenap_5050special_field}
        input_precision=2
        input_id="pasaran_keiasgenap_5050special_field"
        input_placeholder="KEI AS GENAP(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_discasgenap_5050special_field}
        input_precision=2
        input_id="pasaran_discasgenap_5050special_field"
        input_placeholder="DISC AS GENAP(%)"/>
    
    
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keiasbesar_5050special_field}
        input_precision=2
        input_id="pasaran_keiasbesar_5050special_field"
        input_placeholder="KEI AS BESAR(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_discasbesar_5050special_field}
        input_precision=2
        input_id="pasaran_discasbesar_5050special_field"
        input_placeholder="DISC AS BESAR(%)"/>
    
   
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keiaskecil_5050special_field}
        input_precision=2
        input_id="pasaran_keiaskecil_5050special_field"
        input_placeholder="KEI AS KECIL(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_discaskecil_5050special_field}
        input_precision=2
        input_id="pasaran_discaskecil_5050special_field"
        input_placeholder="DISC AS KECIL(%)"/>
    
   
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keikopganjil_5050special_field}
        input_precision=2
        input_id="pasaran_keikopganjil_5050special_field"
        input_placeholder="KEI KOP GANJIL(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_disckopganjil_5050special_field}
        input_precision=2
        input_id="pasaran_disckopganjil_5050special_field"
        input_placeholder="DISC KOP GANJIL(%)"/>
    
    
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keikopgenap_5050special_field}
        input_precision=2
        input_id="pasaran_keikopgenap_5050special_field"
        input_placeholder="KEI KOP GENAP(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_disckopgenap_5050special_field}
        input_precision=2
        input_id="pasaran_disckopgenap_5050special_field"
        input_placeholder="DISC KOP GENAP(%)"/>

    
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keikopbesar_5050special_field}
        input_precision=2
        input_id="pasaran_keikopbesar_5050special_field"
        input_placeholder="KEI KOP BESAR(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_disckopbesar_5050special_field}
        input_precision=2
        input_id="pasaran_disckopbesar_5050special_field"
        input_placeholder="DISC KOP BESAR(%)"/>
    
    
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keikopkecil_5050special_field}
        input_precision=2
        input_id="pasaran_keikopkecil_5050special_field"
        input_placeholder="KEI KOP KECIL(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_disckopkecil_5050special_field}
        input_precision=2
        input_id="pasaran_disckopkecil_5050special_field"
        input_placeholder="DISC KOP KECIL(%)"/>
    
    
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keikepalaganjil_5050special_field}
        input_precision=2
        input_id="pasaran_keikepalaganjil_5050special_field"
        input_placeholder="KEI KEPALA GANJIL(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_disckepalaganjil_5050special_field}
        input_precision=2
        input_id="pasaran_disckepalaganjil_5050special_field"
        input_placeholder="DISC KEPALA GANJIL(%)"/>

    
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keikepalagenap_5050special_field}
        input_precision=2
        input_id="pasaran_keikepalagenap_5050special_field"
        input_placeholder="KEI KEPALA GENAP(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_disckepalagenap_5050special_field}
        input_precision=2
        input_id="pasaran_disckepalagenap_5050special_field"
        input_placeholder="DISC KEPALA GENAP(%)"/>
    
   
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_keikepalabesar_5050special_field}
        input_id="pasaran_keikepalabesar_5050special_field"
        input_placeholder="KEI KEPALA BESAR(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_disckepalabesar_5050special_field}
        input_id="pasaran_disckepalabesar_5050special_field"
        input_placeholder="DISC KEPALA BESAR(%)"/>

    
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_keikepalakecil_5050special_field}
        input_id="pasaran_keikepalakecil_5050special_field"
        input_placeholder="KEI KEPALA KECIL(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_disckepalakecil_5050special_field}
        input_id="pasaran_disckepalakecil_5050special_field"
        input_placeholder="DISC KEPALA KECIL(%)"/>
    
   
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_keiekorganjil_5050special_field}
        input_id="pasaran_keiekorganjil_5050special_field"
        input_placeholder="KEI EKOR GANJIL(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_discekorganjil_5050special_field}
        input_id="pasaran_discekorganjil_5050special_field"
        input_placeholder="DISC EKOR GANJIL(%)"/>
    
   
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_keiekorgenap_5050special_field}
        input_id="pasaran_keiekorgenap_5050special_field"
        input_placeholder="KEI EKOR GENAP(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_discekorgenap_5050special_field}
        input_id="pasaran_discekorgenap_5050special_field"
        input_placeholder="DISC EKOR GENAP(%)"/>
    
   
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_keiekorbesar_5050special_field}
        input_id="pasaran_keiekorbesar_5050special_field"
        input_placeholder="KEI EKOR BESAR(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_discekorbesar_5050special_field}
        input_id="pasaran_discekorbesar_5050special_field"
        input_placeholder="DISC EKOR BESAR(%)"/>
    
   
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_keiekorkecil_5050special_field}
        input_id="pasaran_keiekorkecil_5050special_field"
        input_placeholder="KEI EKOR KECIL(%)"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        input_precision=2
        bind:value={pasaran_discekorkecil_5050special_field}
        input_id="pasaran_discekorkecil_5050special_field"
        input_placeholder="DISC EKOR KECIL(%)"/>
</div>
{#if pasaran_tipe != "WAJIB"}
<button on:click={() => {
    save432d();
}} class="{buttonLoading_class} btn-block">Submit</button>
{/if}