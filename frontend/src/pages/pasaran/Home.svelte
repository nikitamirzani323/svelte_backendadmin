<script>
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Input_custom from '../../components/Input.svelte' 
    import Modal_popup from '../../components/Modal_popup.svelte'
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Panel_432D from '../pasaran/432D.svelte' 
    import Panel_colokbebas from '../pasaran/colok_bebas.svelte' 
    import Panel_colokmacau from '../pasaran/colok_macau.svelte' 
    import Panel_coloknaga from '../pasaran/colok_naga.svelte' 
    import Panel_colokjitu from '../pasaran/colok_jitu.svelte' 
    import Panel_5050umum from '../pasaran/5050_umum.svelte' 
    import Panel_5050special from '../pasaran/5050_special.svelte' 
    import Panel_5050kombinasi from '../pasaran/5050_kombinasi.svelte' 
    import Panel_macaukombinasi from '../pasaran/kombinasi.svelte' 
    import Panel_dasar from '../pasaran/dasar.svelte' 
    import Panel_shio from '../pasaran/shio.svelte' 
    import Loader from '../../components/Loader.svelte' 
    import Panel from '../../components/Panel_default.svelte' 

    export let path_api = "";
    export let token = "";
    export let listHome = [];
    export let totalrecord = 0;

    let page = "Pasaran";
    let sData = "New";
    let isModal_Form_New = false
    let isModal_Form_pasaran = false
    let isModalLoading = false
    let isModalNotif = false
    let modal_width = "max-w-xl"
    let modalpasaran_width = "max-w-xl"
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let buttonLoading2_class = "btn btn-primary"
    let msg_error = "";
    let searchHome = "";
    let filterHome = [];

    let idcomppasaran = "";
    let pasaran_tipe = "";
    let select_pasaranonline = "";
    let pasaran_idpasarantogel_field = "";
	let pasaran_name_field = "";
	let pasaran_tipepasaran_field = "";
	let pasaran_diundi_field = "";
	let pasaran_situs_field = "";
	let pasaran_tutup_field = "";
	let pasaran_jadwal_field = "";
	let pasaran_buka_field = "";
    let pasaran_display_field = 0;
	let pasaran_status_field = "";
	let pasaran_create_field = "";
	let pasaran_createdate_field = "";
	let pasaran_update_field = "";
	let pasaran_updatedate_field = "";

	let pasaran_limitline4d_field = 0;
	let pasaran_limitline3d_field = 0;
	let pasaran_limitline3dd_field = 0;
	let pasaran_limitline2d_field = 0;
	let pasaran_limitline2dd_field = 0;
	let pasaran_limitline2dt_field = 0;
	let pasaran_bbfs_field = 0;
	let pasaran_minbet_432d_field = 0;
	let pasaran_maxbet4d_432d_field = 0;
	let pasaran_maxbet3d_432d_field = 0;
	let pasaran_maxbet3dd_432d_field = 0;
	let pasaran_maxbet2d_432d_field = 0;
	let pasaran_maxbet2dd_432d_field = 0;
	let pasaran_maxbet2dt_432d_field = 0;
	let pasaran_limitotal4d_432d_field = 0;
	let pasaran_limitotal3d_432d_field = 0;
	let pasaran_limitotal3dd_432d_field = 0;
	let pasaran_limitotal2d_432d_field = 0;
	let pasaran_limitotal2dd_432d_field = 0;
	let pasaran_limitotal2dt_432d_field = 0;
	let pasaran_limitglobal4d_432d_field = 0;
	let pasaran_limitglobal3d_432d_field = 0;
	let pasaran_limitglobal3dd_432d_field = 0;
	let pasaran_limitglobal2d_432d_field = 0;
	let pasaran_limitglobal2dd_432d_field = 0;
	let pasaran_limitglobal2dt_432d_field = 0;
	let pasaran_disc4d_432d_field = 0;
	let pasaran_disc3d_432d_field = 0;
	let pasaran_disc3dd_432d_field = 0;
	let pasaran_disc2d_432d_field = 0;
	let pasaran_disc2dd_432d_field = 0;
	let pasaran_disc2dt_432d_field = 0;
	let pasaran_win4d_432d_field = 0;
	let pasaran_win3d_432d_field = 0;
	let pasaran_win3dd_432d_field = 0;
	let pasaran_win2d_432d_field = 0;
	let pasaran_win2dd_432d_field = 0;
	let pasaran_win2dt_432d_field = 0;
	let pasaran_win4dnodisc_432d_field = 0;
	let pasaran_win3dnodisc_432d_field = 0;
	let pasaran_win3ddnodisc_432d_field = 0;
	let pasaran_win2dnodisc_432d_field = 0;
	let pasaran_win2ddnodisc_432d_field = 0;
	let pasaran_win2dtnodisc_432d_field = 0;
	let pasaran_win4d_bb_kena_432d_field = 0;
    let pasaran_win3d_bb_kena_432d_field = 0;
    let pasaran_win3dd_bb_kena_432d_field = 0;
    let pasaran_win2d_bb_kena_432d_field = 0;
    let pasaran_win2dd_bb_kena_432d_field = 0;
    let pasaran_win2dt_bb_kena_432d_field = 0;
    let pasaran_win4d_bb_432d_field = 0;
    let pasaran_win3d_bb_432d_field = 0;
    let pasaran_win3dd_bb_432d_field = 0;
    let pasaran_win2d_bb_432d_field = 0;
    let pasaran_win2dd_bb_432d_field = 0;
    let pasaran_win2dt_bb_432d_field = 0;
	let pasaran_minbet_cbebas_field = 0;
	let pasaran_maxbet_cbebas_field = 0;
	let pasaran_limitotal_cbebas_field = 0;
	let pasaran_limitglobal_cbebas_field = 0;
	let pasaran_win_cbebas_field = 0;
	let pasaran_disc_cbebas_field = 0;
	let pasaran_minbet_cmacau_field = 0;
	let pasaran_maxbet_cmacau_field = 0;
	let pasaran_limitotal_cmacau_field = 0;
	let pasaran_limitglobal_cmacau_field = 0;
	let pasaran_win2_cmacau_field = 0;
	let pasaran_win3_cmacau_field = 0;
	let pasaran_win4_cmacau_field = 0;
	let pasaran_disc_cmacau_field = 0;
	let pasaran_minbet_cnaga_field = 0;
	let pasaran_maxbet_cnaga_field = 0;
	let pasaran_win3_cnaga_field = 0;
	let pasaran_win4_cnaga_field = 0;
	let pasaran_disc_cnaga_field = 0;
	let pasaran_limitglobal_cnaga_field = 0;
	let pasaran_limittotal_cnaga_field = 0;
	let pasaran_minbet_cjitu_field = 0;
	let pasaran_maxbet_cjitu_field = 0;
	let pasaran_winas_cjitu_field = 0;
	let pasaran_winkop_cjitu_field = 0;
	let pasaran_winkepala_cjitu_field = 0;
	let pasaran_winekor_cjitu_field = 0;
	let pasaran_desc_cjitu_field = 0;
	let pasaran_limitglobal_cjitu_field = 0;
	let pasaran_limittotal_cjitu_field = 0;
	let pasaran_minbet_5050umum_field = 0;
	let pasaran_maxbet_5050umum_field = 0;
	let pasaran_keibesar_5050umum_field = 0;
	let pasaran_keikecil_5050umum_field = 0;
	let pasaran_keigenap_5050umum_field = 0;
	let pasaran_keiganjil_5050umum_field = 0;
	let pasaran_keitengah_5050umum_field = 0;
	let pasaran_keitepi_5050umum_field = 0;
	let pasaran_discbesar_5050umum_field = 0;
	let pasaran_disckecil_5050umum_field = 0;
	let pasaran_discgenap_5050umum_field = 0;
	let pasaran_discganjil_5050umum_field = 0;
	let pasaran_disctengah_5050umum_field = 0;
	let pasaran_disctepi_5050umum_field = 0;
	let pasaran_limitglobal_5050umum_field = 0;
	let pasaran_limittotal_5050umum_field = 0;
	let pasaran_minbet_5050special_field = 0;
	let pasaran_maxbet_5050special_field = 0;
	let pasaran_keiasganjil_5050special_field = 0;
	let pasaran_keiasgenap_5050special_field = 0;
	let pasaran_keiasbesar_5050special_field = 0;
	let pasaran_keiaskecil_5050special_field = 0;
	let pasaran_keikopganjil_5050special_field = 0;
	let pasaran_keikopgenap_5050special_field = 0;
	let pasaran_keikopbesar_5050special_field = 0;
	let pasaran_keikopkecil_5050special_field = 0;
	let pasaran_keikepalaganjil_5050special_field = 0;
	let pasaran_keikepalagenap_5050special_field = 0;
	let pasaran_keikepalabesar_5050special_field = 0;
	let pasaran_keikepalakecil_5050special_field = 0;
	let pasaran_keiekorganjil_5050special_field = 0;
	let pasaran_keiekorgenap_5050special_field = 0;
	let pasaran_keiekorbesar_5050special_field = 0;
	let pasaran_keiekorkecil_5050special_field = 0;
	let pasaran_discasganjil_5050special_field = 0;
	let pasaran_discasgenap_5050special_field = 0;
	let pasaran_discasbesar_5050special_field = 0;
	let pasaran_discaskecil_5050special_field = 0;
	let pasaran_disckopganjil_5050special_field = 0;
	let pasaran_disckopgenap_5050special_field = 0;
	let pasaran_disckopbesar_5050special_field = 0;
	let pasaran_disckopkecil_5050special_field = 0;
	let pasaran_disckepalaganjil_5050special_field = 0;
	let pasaran_disckepalagenap_5050special_field = 0;
	let pasaran_disckepalabesar_5050special_field = 0;
	let pasaran_disckepalakecil_5050special_field = 0;
	let pasaran_discekorganjil_5050special_field = 0;
	let pasaran_discekorgenap_5050special_field = 0;
	let pasaran_discekorbesar_5050special_field = 0;
	let pasaran_discekorkecil_5050special_field = 0;
	let pasaran_limitglobal_5050special_field = 0;
	let pasaran_limittotal_5050special_field = 0;
	let pasaran_minbet_5050kombinasi_field = 0;
	let pasaran_maxbet_5050kombinasi_field = 0;
	let pasaran_belakangkeimono_5050kombinasi_field = 0;
	let pasaran_belakangkeistereo_5050kombinasi_field = 0;
	let pasaran_belakangkeikembang_5050kombinasi_field = 0;
	let pasaran_belakangkeikempis_5050kombinasi_field = 0;
	let pasaran_belakangkeikembar_5050kombinasi_field = 0;
	let pasaran_tengahkeimono_5050kombinasi_field = 0;
	let pasaran_tengahkeistereo_5050kombinasi_field = 0;
	let pasaran_tengahkeikembang_5050kombinasi_field = 0;
	let pasaran_tengahkeikempis_5050kombinasi_field = 0;
	let pasaran_tengahkeikembar_5050kombinasi_field = 0;
	let pasaran_depankeimono_5050kombinasi_field = 0;
	let pasaran_depankeistereo_5050kombinasi_field = 0;
	let pasaran_depankeikembang_5050kombinasi_field = 0;
	let pasaran_depankeikempis_5050kombinasi_field = 0;
	let pasaran_depankeikembar_5050kombinasi_field = 0;
	let pasaran_belakangdiscmono_5050kombinasi_field = 0;
	let pasaran_belakangdiscstereo_5050kombinasi_field = 0;
	let pasaran_belakangdisckembang_5050kombinasi_field = 0;
	let pasaran_belakangdisckempis_5050kombinasi_field = 0;
	let pasaran_belakangdisckembar_5050kombinasi_field = 0;
	let pasaran_tengahdiscmono_5050kombinasi_field = 0;
	let pasaran_tengahdiscstereo_5050kombinasi_field = 0;
	let pasaran_tengahdisckembang_5050kombinasi_field = 0;
	let pasaran_tengahdisckempis_5050kombinasi_field = 0;
	let pasaran_tengahdisckembar_5050kombinasi_field = 0;
	let pasaran_depandiscmono_5050kombinasi_field = 0;
	let pasaran_depandiscstereo_5050kombinasi_field = 0;
	let pasaran_depandisckembang_5050kombinasi_field = 0;
	let pasaran_depandisckempis_5050kombinasi_field = 0;
	let pasaran_depandisckembar_5050kombinasi_field = 0;
	let pasaran_limitglobal_5050kombinasi_field = 0;
	let pasaran_limittotal_5050kombinasi_field = 0;
	let pasaran_minbet_kombinasi_field = 0;
	let pasaran_maxbet_kombinasi_field = 0;
	let pasaran_win_kombinasi_field = 0;
	let pasaran_disc_kombinasi_field = 0;
	let pasaran_limitglobal_kombinasi_field = 0;
	let pasaran_limittotal_kombinasi_field = 0;
	let pasaran_minbet_dasar_field = 0;
	let pasaran_maxbet_dasar_field = 0;
	let pasaran_keibesar_dasar_field = 0;
	let pasaran_keikecil_dasar_field = 0;
	let pasaran_keigenap_dasar_field = 0;
	let pasaran_keiganjil_dasar_field = 0;
	let pasaran_discbesar_dasar_field = 0;
	let pasaran_disckecil_dasar_field = 0;
	let pasaran_discgenap_dasar_field = 0;
	let pasaran_discganjil_dasar_field = 0;
	let pasaran_limitglobal_dasar_field = 0;
	let pasaran_limittotal_dasar_field = 0;
	let pasaran_minbet_shio_field = 0;
	let pasaran_maxbet_shio_field = 0;
	let pasaran_win_shio_field = 0;
	let pasaran_disc_shio_field = 0;
	let pasaran_shioyear_shio_field = "";
	let pasaran_limitglobal_shio_field = 0;
	let pasaran_limittotal_shio_field = 0;
	let listPasaranOnline = [];

    let permainan = "";
    let tab_pasaranonline = "bg-sky-600 text-white"
    let tab_bbfs = ""
    let tab_configure = ""
    let panel_pasaranonline = true
    let panel_bbfs = false
    let panel_configure = false

    let panel_432D = false
    let panel_cbebas = false
    let panel_cmacau = false
    let panel_cnaga = false
    let panel_cjitu = false
    let panel_5050umum = false
    let panel_5050special = false
    let panel_5050kombinasi = false
    let panel_macaukombinasi = false
    let panel_dasar = false
    let panel_shio = false

    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        form_pasaran_situs_field: yup
            .string()
            .required("Situs is Required"),
        form_pasaran_diundi_field: yup
            .string()
            .required("Diundi is Required")
            .matches(
                /^[a-zA-z0-9 ]+$/,
                "Diundi must Character A-Z or a-z or 1-9"
            )
            .min(4, "Diundi must be at least 4 Character")
            .max(70, "Diundi must be at most 70 Character"),
        form_pasaran_tutup_field: yup
            .string()
            .required("Jam Tutup is Required")
            .matches(
                /^(?:[01]\d|2[0123]):(?:[012345]\d)+$/,
                "Jam Tutup example 24:24 or 12:30"
            ),
        form_pasaran_jadwal_field: yup
            .string()
            .required("Jam Jadwal is Required")
            .matches(
                /^(?:[01]\d|2[0123]):(?:[012345]\d)+$/,
                "Jam Jadwal example 24:24 or 12:30"
            ),
        form_pasaran_open_field: yup
            .string()
            .required("Jam Open is Required")
            .matches(
                /^(?:[01]\d|2[0123]):(?:[012345]\d)+$/,
                "Jam Open example 24:24 or 12:30"
            ),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            form_pasaran_situs_field: "",
            form_pasaran_diundi_field: "",
            form_pasaran_tutup_field: "",
            form_pasaran_jadwal_field: "",
            form_pasaran_open_field: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(
                values.form_pasaran_situs_field,
                values.form_pasaran_diundi_field,
                values.form_pasaran_tutup_field,
                values.form_pasaran_jadwal_field,
                values.form_pasaran_open_field);
        },
    });
    async function SaveTransaksi(situs,diundi,jamtutup,jamjadwal,jamopen) {
        let flag = true;
        msg_error = "";
        
        if (flag) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savepasaran", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page: "ADMINRULE-SAVE",
                    idpasaran: idcomppasaran,
                    pasaran_diundi: diundi,
                    pasaran_url: situs,
                    pasaran_jamtutup: jamtutup,
                    pasaran_jamjadwal: jamjadwal,
                    pasaran_jamopen: jamopen,
                    pasaran_display: parseInt(pasaran_display_field),
                    pasaran_status: pasaran_status_field,
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
    
    async function EditData(e,y) {
        if(e != ""){
            isModalLoading = true;
            modal_width = "max-w-6xl";
            sData = "Edit";
            idcomppasaran = e;
            pasaran_tipe = y;
            clearField()
            const res = await fetch(path_api+"api/editpasaran", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                }),
            });
            const json = await res.json();
            let record = json.record;
            let record_pasaranonline = json.pasaranonline;
            if (json.status === 400) {
                dispatch("handleLogout", "call");
            }else if(json.status === 200) {
                isModal_Form_New = true;
                isModalLoading = false;
                listPasaranOnline = [];
                for (let i = 0; i < record.length; i++) {
                    pasaran_idpasarantogel_field = record[i]["idpasarantogel"];
                    pasaran_name_field = record[i]["nmpasaran"];
                    pasaran_tipepasaran_field = y;
                    $form.form_pasaran_situs_field = record[i]["pasaranurl"];
                    $form.form_pasaran_diundi_field = record[i]["pasarandiundi"];
                    $form.form_pasaran_tutup_field = record[i]["jamtutup"];
                    $form.form_pasaran_jadwal_field = record[i]["jamjadwal"];
                    $form.form_pasaran_open_field = record[i]["jamopen"];
                    pasaran_display_field = record[i]["displaypasaran"];
                    pasaran_status_field = record[i]["statuspasaranactive"];
                    pasaran_create_field = record[i]["create"];
                    pasaran_createdate_field = record[i]["createdate"];
                    pasaran_update_field = record[i]["update"];
                    pasaran_updatedate_field = record[i]["updatedate"];

                    pasaran_limitline4d_field = record[i]["limitline_4d"];
                    pasaran_limitline3d_field = record[i]["limitline_3d"];
                    pasaran_limitline3dd_field = record[i]["limitline_3dd"];
                    pasaran_limitline2d_field = record[i]["limitline_2d"];
                    pasaran_limitline2dd_field = record[i]["limitline_2dd"];
                    pasaran_limitline2dt_field = record[i]["limitline_2dt"];
                    pasaran_bbfs_field = record[i]["bbfs"];
                    pasaran_minbet_432d_field = record[i]["minbet_432d"];
                    pasaran_maxbet4d_432d_field = record[i]["maxbet4d_432d"];
                    pasaran_maxbet3d_432d_field = record[i]["maxbet3d_432d"];
                    pasaran_maxbet3dd_432d_field = record[i]["maxbet3dd_432d"];
                    pasaran_maxbet2d_432d_field = record[i]["maxbet2d_432d"];
                    pasaran_maxbet2dd_432d_field = record[i]["maxbet2dd_432d"];
                    pasaran_maxbet2dt_432d_field = record[i]["maxbet2dt_432d"];
                    pasaran_limitotal4d_432d_field = record[i]["limitotal4d_432d"];
                    pasaran_limitotal3d_432d_field = record[i]["limitotal3d_432d"];
                    pasaran_limitotal3dd_432d_field = record[i]["limitotal3dd_432d"];
                    pasaran_limitotal2d_432d_field = record[i]["limitotal2d_432d"];
                    pasaran_limitotal2dd_432d_field = record[i]["limitotal2dd_432d"];
                    pasaran_limitotal2dt_432d_field = record[i]["limitotal2dt_432d"];
                    pasaran_limitglobal4d_432d_field = record[i]["limitglobal4d_432d"];
                    pasaran_limitglobal3d_432d_field = record[i]["limitglobal3d_432d"];
                    pasaran_limitglobal3dd_432d_field = record[i]["limitglobal3dd_432d"];
                    pasaran_limitglobal2d_432d_field = record[i]["limitglobal2d_432d"];
                    pasaran_limitglobal2dd_432d_field = record[i]["limitglobal2dd_432d"];
                    pasaran_limitglobal2dt_432d_field = record[i]["limitglobal2dt_432d"];
                    pasaran_disc4d_432d_field = (record[i]["disc4d_432d"] * 100).toFixed(2);
                    pasaran_disc3d_432d_field = (record[i]["disc3d_432d"] * 100).toFixed(2);
                    pasaran_disc3dd_432d_field = (record[i]["disc3dd_432d"] * 100).toFixed(2);
                    pasaran_disc2d_432d_field = (record[i]["disc2d_432d"] * 100).toFixed(2);
                    pasaran_disc2dd_432d_field = (record[i]["disc2dd_432d"] * 100).toFixed(2);
                    pasaran_disc2dt_432d_field = (record[i]["disc2dt_432d"] * 100).toFixed(2);
                    pasaran_win4d_432d_field = record[i]["win4d_432d"];
                    pasaran_win3d_432d_field = record[i]["win3d_432d"];
                    pasaran_win3dd_432d_field = record[i]["win3dd_432d"];
                    pasaran_win2d_432d_field = record[i]["win2d_432d"];
                    pasaran_win2dd_432d_field = record[i]["win2dd_432d"];
                    pasaran_win2dt_432d_field = record[i]["win2dt_432d"];
                    pasaran_win4dnodisc_432d_field = record[i]["win4dnodisc_432d"];
                    pasaran_win3dnodisc_432d_field = record[i]["win3dnodisc_432d"];
                    pasaran_win3ddnodisc_432d_field = record[i]["win3ddnodisc_432d"];
                    pasaran_win2dnodisc_432d_field = record[i]["win2dnodisc_432d"];
                    pasaran_win2ddnodisc_432d_field = record[i]["win2ddnodisc_432d"];
                    pasaran_win2dtnodisc_432d_field = record[i]["win2dtnodisc_432d"];
                    pasaran_win4d_bb_kena_432d_field = record[i]["win4dbb_kena_432d"];
                    pasaran_win3d_bb_kena_432d_field = record[i]["win3dbb_kena_432d"];
                    pasaran_win3dd_bb_kena_432d_field = record[i]["win3ddbb_kena_432d"];
                    pasaran_win2d_bb_kena_432d_field = record[i]["win2dbb_kena_432d"];
                    pasaran_win2dd_bb_kena_432d_field = record[i]["win2ddbb_kena_432d"];
                    pasaran_win2dt_bb_kena_432d_field = record[i]["win2dtbb_kena_432d"];
                    pasaran_win4d_bb_432d_field = record[i]["win4dbb_432d"];
                    pasaran_win3d_bb_432d_field = record[i]["win3dbb_432d"];
                    pasaran_win3dd_bb_432d_field = record[i]["win3ddbb_432d"];
                    pasaran_win2d_bb_432d_field = record[i]["win2dbb_432d"];
                    pasaran_win2dd_bb_432d_field = record[i]["win2ddbb_432d"];
                    pasaran_win2dt_bb_432d_field = record[i]["win2dtbb_432d"];

                    pasaran_minbet_cbebas_field = record[i]["minbet_cbebas"];
                    pasaran_maxbet_cbebas_field = record[i]["maxbet_cbebas"];
                    pasaran_limitotal_cbebas_field = record[i]["limittotal_cbebas"];
                    pasaran_limitglobal_cbebas_field = record[i]["limitglobal_cbebas"];
                    pasaran_win_cbebas_field = record[i]["win_cbebas"].toFixed(2);
                    pasaran_disc_cbebas_field = (record[i]["disc_cbebas"] * 100).toFixed(2);
                    pasaran_minbet_cmacau_field = record[i]["minbet_cmacau"];
                    pasaran_maxbet_cmacau_field = record[i]["maxbet_cmacau"];
                    pasaran_limitotal_cmacau_field = record[i]["limitotal_cmacau"];
                    pasaran_limitglobal_cmacau_field = record[i]["limitglobal_cmacau"];
                    pasaran_win2_cmacau_field = record[i]["win2d_cmacau"].toFixed(2);
                    pasaran_win3_cmacau_field = record[i]["win3d_cmacau"].toFixed(2);
                    pasaran_win4_cmacau_field = record[i]["win4d_cmacau"].toFixed(2);
                    pasaran_disc_cmacau_field = (record[i]["disc_cmacau"] * 100).toFixed(2);
                    pasaran_minbet_cnaga_field = record[i]["minbet_cnaga"];
                    pasaran_maxbet_cnaga_field = record[i]["maxbet_cnaga"];
                    pasaran_win3_cnaga_field = record[i]["win3_cnaga"].toFixed(2);
                    pasaran_win4_cnaga_field = record[i]["win4_cnaga"].toFixed(2);
                    pasaran_disc_cnaga_field = (record[i]["disc_cnaga"] * 100).toFixed(2);
                    pasaran_limitglobal_cnaga_field = record[i]["limitglobal_cnaga"];
                    pasaran_limittotal_cnaga_field = record[i]["limittotal_cnaga"];
                    pasaran_minbet_cjitu_field = record[i]["minbet_cjitu"];
                    pasaran_maxbet_cjitu_field = record[i]["maxbet_cjitu"];
                    pasaran_winas_cjitu_field = record[i]["winas_cjitu"].toFixed(2);
                    pasaran_winkop_cjitu_field = record[i]["winkop_cjitu"].toFixed(2);
                    pasaran_winkepala_cjitu_field = record[i]["winkepala_cjitu"].toFixed(2);
                    pasaran_winekor_cjitu_field = record[i]["winekor_cjitu"].toFixed(2);
                    pasaran_desc_cjitu_field = (record[i]["desc_cjitu"] * 100).toFixed(2);
                    pasaran_limitglobal_cjitu_field = record[i]["limitglobal_cjitu"];
                    pasaran_limittotal_cjitu_field = record[i]["limittotal_cjitu"];
                    pasaran_minbet_5050umum_field = record[i]["minbet_5050umum"];
                    pasaran_maxbet_5050umum_field = record[i]["maxbet_5050umum"];
                    pasaran_keibesar_5050umum_field = (record[i]["keibesar_5050umum"] * 100).toFixed(2);
                    pasaran_keikecil_5050umum_field = (record[i]["keikecil_5050umum"] * 100).toFixed(2);
                    pasaran_keigenap_5050umum_field = (record[i]["keigenap_5050umum"] * 100).toFixed(2);
                    pasaran_keiganjil_5050umum_field = (record[i]["keiganjil_5050umum"] * 100).toFixed(2);
                    pasaran_keitengah_5050umum_field = (record[i]["keitengah_5050umum"] * 100).toFixed(2);
                    pasaran_keitepi_5050umum_field = (record[i]["keitepi_5050umum"] * 100).toFixed(2);
                    pasaran_discbesar_5050umum_field = (record[i]["discbesar_5050umum"] * 100).toFixed(2);
                    pasaran_disckecil_5050umum_field = (record[i]["disckecil_5050umum"] * 100).toFixed(2);
                    pasaran_discgenap_5050umum_field = (record[i]["discgenap_5050umum"] * 100).toFixed(2);
                    pasaran_discganjil_5050umum_field = (record[i]["discganjil_5050umum"] * 100).toFixed(2);
                    pasaran_disctengah_5050umum_field = (record[i]["disctengah_5050umum"] * 100).toFixed(2);
                    pasaran_disctepi_5050umum_field = (record[i]["disctepi_5050umum"] * 100).toFixed(2);
                    pasaran_limitglobal_5050umum_field = record[i]["limitglobal_5050umum"];
                    pasaran_limittotal_5050umum_field = record[i]["limittotal_5050umum"];
                    pasaran_minbet_5050special_field = record[i]["minbet_5050special"];
                    pasaran_maxbet_5050special_field = record[i]["maxbet_5050special"];
                    pasaran_keiasganjil_5050special_field = (record[i]["keiasganjil_5050special"] * 100).toFixed(2);
                    pasaran_keiasgenap_5050special_field = (record[i]["keiasgenap_5050special"] * 100).toFixed(2);
                    pasaran_keiasbesar_5050special_field = (record[i]["keiasbesar_5050special"] * 100).toFixed(2);
                    pasaran_keiaskecil_5050special_field = (record[i]["keiaskecil_5050special"] * 100).toFixed(2);
                    pasaran_keikopganjil_5050special_field = (record[i]["keikopganjil_5050special"] * 100).toFixed(2);
                    pasaran_keikopgenap_5050special_field = (record[i]["keikopgenap_5050special"] * 100).toFixed(2);
                    pasaran_keikopbesar_5050special_field = (record[i]["keikopbesar_5050special"] * 100).toFixed(2);
                    pasaran_keikopkecil_5050special_field = (record[i]["keikopkecil_5050special"] * 100).toFixed(2);
                    pasaran_keikepalaganjil_5050special_field = (record[i]["keikepalaganjil_5050special"] * 100).toFixed(2);
                    pasaran_keikepalagenap_5050special_field = (record[i]["keikepalagenap_5050special"] * 100).toFixed(2);
                    pasaran_keikepalabesar_5050special_field = (record[i]["keikepalabesar_5050special"] * 100).toFixed(2);
                    pasaran_keikepalakecil_5050special_field = (record[i]["keikepalakecil_5050special"] * 100).toFixed(2);
                    pasaran_keiekorganjil_5050special_field = (record[i]["keiekorganjil_5050special"] * 100).toFixed(2);
                    pasaran_keiekorgenap_5050special_field = (record[i]["keiekorgenap_5050special"] * 100).toFixed(2);
                    pasaran_keiekorbesar_5050special_field = (record[i]["keiekorbesar_5050special"] * 100).toFixed(2);
                    pasaran_keiekorkecil_5050special_field = (record[i]["keiekorkecil_5050special"] * 100).toFixed(2);
                    pasaran_discasganjil_5050special_field = (record[i]["discasganjil_5050special"] * 100).toFixed(2);
                    pasaran_discasgenap_5050special_field = (record[i]["discasgenap_5050special"] * 100).toFixed(2);
                    pasaran_discasbesar_5050special_field = (record[i]["discasbesar_5050special"] * 100).toFixed(2);
                    pasaran_discaskecil_5050special_field = (record[i]["discaskecil_5050special"] * 100).toFixed(2);
                    pasaran_disckopganjil_5050special_field = (record[i]["disckopganjil_5050special"] * 100).toFixed(2);
                    pasaran_disckopgenap_5050special_field = (record[i]["disckopgenap_5050special"] * 100).toFixed(2);
                    pasaran_disckopbesar_5050special_field = (record[i]["disckopbesar_5050special"] * 100).toFixed(2);
                    pasaran_disckopkecil_5050special_field = (record[i]["disckopkecil_5050special"] * 100).toFixed(2);
                    pasaran_disckepalaganjil_5050special_field = (record[i]["disckepalaganjil_5050special"] * 100).toFixed(2);
                    pasaran_disckepalagenap_5050special_field = (record[i]["disckepalagenap_5050special"] * 100).toFixed(2);
                    pasaran_disckepalabesar_5050special_field = (record[i]["disckepalabesar_5050special"] * 100).toFixed(2);
                    pasaran_disckepalakecil_5050special_field = (record[i]["disckepalakecil_5050special"] * 100).toFixed(2);
                    pasaran_discekorganjil_5050special_field = (record[i]["discekorganjil_5050special"] * 100).toFixed(2);
                    pasaran_discekorgenap_5050special_field = (record[i]["discekorgenap_5050special"] * 100).toFixed(2);
                    pasaran_discekorbesar_5050special_field = (record[i]["discekorbesar_5050special"] * 100).toFixed(2);
                    pasaran_discekorkecil_5050special_field = (record[i]["discekorkecil_5050special"] * 100).toFixed(2);
                    pasaran_limitglobal_5050special_field = record[i]["limitglobal_5050special"];
                    pasaran_limittotal_5050special_field = record[i]["limittotal_5050special"];
                    pasaran_minbet_5050kombinasi_field = record[i]["minbet_5050kombinasi"];
                    pasaran_maxbet_5050kombinasi_field = record[i]["maxbet_5050kombinasi"];
                    pasaran_belakangkeimono_5050kombinasi_field = (record[i]["belakangkeimono_5050kombinasi"] * 100).toFixed(2);
                    pasaran_belakangkeistereo_5050kombinasi_field = (record[i]["belakangkeistereo_5050kombinasi"] * 100).toFixed(2);
                    pasaran_belakangkeikembang_5050kombinasi_field = (record[i]["belakangkeikembang_5050kombinasi"] * 100).toFixed(2);
                    pasaran_belakangkeikempis_5050kombinasi_field = (record[i]["belakangkeikempis_5050kombinasi"] * 100).toFixed(2);
                    pasaran_belakangkeikembar_5050kombinasi_field = (record[i]["belakangkeikembar_5050kombinasi"] * 100).toFixed(2);
                    pasaran_tengahkeimono_5050kombinasi_field = (record[i]["tengahkeimono_5050kombinasi"] * 100).toFixed(2);
                    pasaran_tengahkeistereo_5050kombinasi_field = (record[i]["tengahkeistereo_5050kombinasi"] * 100).toFixed(2);
                    pasaran_tengahkeikembang_5050kombinasi_field = (record[i]["tengahkeikembang_5050kombinasi"] * 100).toFixed(2);
                    pasaran_tengahkeikempis_5050kombinasi_field = (record[i]["tengahkeikempis_5050kombinasi"] * 100).toFixed(2);
                    pasaran_tengahkeikembar_5050kombinasi_field = (record[i]["tengahkeikembar_5050kombinasi"] * 100).toFixed(2);
                    pasaran_depankeimono_5050kombinasi_field = (record[i]["depankeimono_5050kombinasi"] * 100).toFixed(2);
                    pasaran_depankeistereo_5050kombinasi_field = (record[i]["depankeistereo_5050kombinasi"] * 100).toFixed(2);
                    pasaran_depankeikembang_5050kombinasi_field = (record[i]["depankeikembang_5050kombinasi"] * 100).toFixed(2);
                    pasaran_depankeikempis_5050kombinasi_field = (record[i]["depankeikempis_5050kombinasi"] * 100).toFixed(2);
                    pasaran_depankeikembar_5050kombinasi_field = (record[i]["depankeikembar_5050kombinasi"] * 100).toFixed(2);
                    pasaran_belakangdiscmono_5050kombinasi_field = (record[i]["belakangdiscmono_5050kombinasi"] * 100).toFixed(2);
                    pasaran_belakangdiscstereo_5050kombinasi_field = (record[i]["belakangdiscstereo_5050kombinasi"] * 100).toFixed(2);
                    pasaran_belakangdisckembang_5050kombinasi_field = (record[i]["belakangdisckembang_5050kombinasi"] * 100).toFixed(2);
                    pasaran_belakangdisckempis_5050kombinasi_field = (record[i]["belakangdisckempis_5050kombinasi"] * 100).toFixed(2);
                    pasaran_belakangdisckembar_5050kombinasi_field = (record[i]["belakangdisckembar_5050kombinasi"] * 100).toFixed(2);
                    pasaran_tengahdiscmono_5050kombinasi_field = (record[i]["tengahdiscmono_5050kombinasi"] * 100).toFixed(2);
                    pasaran_tengahdiscstereo_5050kombinasi_field = (record[i]["tengahdiscstereo_5050kombinasi"] * 100).toFixed(2);
                    pasaran_tengahdisckembang_5050kombinasi_field = (record[i]["tengahdisckembang_5050kombinasi"] * 100).toFixed(2);
                    pasaran_tengahdisckempis_5050kombinasi_field = (record[i]["tengahdisckempis_5050kombinasi"] * 100).toFixed(2);
                    pasaran_tengahdisckembar_5050kombinasi_field = (record[i]["tengahdisckembar_5050kombinasi"] * 100).toFixed(2);
                    pasaran_depandiscmono_5050kombinasi_field = (record[i]["depandiscmono_5050kombinasi"] * 100).toFixed(2);
                    pasaran_depandiscstereo_5050kombinasi_field = (record[i]["depandiscstereo_5050kombinasi"] * 100).toFixed(2);
                    pasaran_depandisckembang_5050kombinasi_field = (record[i]["depandisckembang_5050kombinasi"] * 100).toFixed(2);
                    pasaran_depandisckempis_5050kombinasi_field = (record[i]["depandisckempis_5050kombinasi"] * 100).toFixed(2);
                    pasaran_depandisckembar_5050kombinasi_field = (record[i]["depandisckembar_5050kombinasi"] * 100).toFixed(2);
                    pasaran_limitglobal_5050kombinasi_field = record[i]["limitglobal_5050kombinasi"];
                    pasaran_limittotal_5050kombinasi_field = record[i]["limittotal_5050kombinasi"];
                    pasaran_minbet_kombinasi_field = record[i]["minbet_kombinasi"];
                    pasaran_maxbet_kombinasi_field = record[i]["maxbet_kombinasi"];
                    pasaran_win_kombinasi_field = record[i]["win_kombinasi"].toFixed(2);
                    pasaran_disc_kombinasi_field = (record[i]["disc_kombinasi"] * 100).toFixed(2);
                    pasaran_limitglobal_kombinasi_field = record[i]["limitglobal_kombinasi"];
                    pasaran_limittotal_kombinasi_field = record[i]["limittotal_kombinasi"];
                    pasaran_minbet_dasar_field = record[i]["minbet_dasar"];
                    pasaran_maxbet_dasar_field = record[i]["maxbet_dasar"];
                    pasaran_keibesar_dasar_field = (record[i]["keibesar_dasar"] * 100).toFixed(2);
                    pasaran_keikecil_dasar_field = (record[i]["keikecil_dasar"] * 100).toFixed(2);
                    pasaran_keigenap_dasar_field = (record[i]["keigenap_dasar"] * 100).toFixed(2);
                    pasaran_keiganjil_dasar_field = (record[i]["keiganjil_dasar"] * 100).toFixed(2);
                    pasaran_discbesar_dasar_field = (record[i]["discbesar_dasar"] * 100).toFixed(2);
                    pasaran_disckecil_dasar_field = (record[i]["disckecil_dasar"] * 100).toFixed(2);
                    pasaran_discgenap_dasar_field = (record[i]["discgenap_dasar"] * 100).toFixed(2);
                    pasaran_discganjil_dasar_field = (record[i]["discganjil_dasar"] * 100).toFixed(2);
                    pasaran_limitglobal_dasar_field = record[i]["limitglobal_dasar"];
                    pasaran_limittotal_dasar_field = record[i]["limittotal_dasar"];
                    
                    pasaran_minbet_shio_field = record[i]["minbet_shio"];
                    pasaran_maxbet_shio_field = record[i]["maxbet_shio"];
                    pasaran_win_shio_field = record[i]["win_shio"].toFixed(2);
                    pasaran_disc_shio_field = (record[i]["disc_shio"] * 100).toFixed(2);
                    pasaran_shioyear_shio_field = record[i]["shioyear_shio"];
                    pasaran_limitglobal_shio_field = record[i]["limitglobal_shio"];
                    pasaran_limittotal_shio_field = record[i]["limittotal_shio"];
                }
                for (var i = 0; i < record_pasaranonline.length; i++) {
                    listPasaranOnline = [
                        ...listPasaranOnline,
                        {
                            idpasaranonline: record_pasaranonline[i]["Idpasaranonline"],
                            haripasaran: record_pasaranonline[i]["haripasaran"],
                        },
                    ];
                }
            }else{
                idcomppasaran = "";
                isModalLoading = false;
                isModalNotif = true;
                msg_error = "Silahkan Hubungi Administrator"
            }
        }
    }
    async function savePasaranOnline() {
        let flag = false;
        msg_error = "";
        if (select_pasaranonline == "" || select_pasaranonline == null) {
            flag = true;
            msg_error += "The Pasaran Online is required";
        }
        if (flag == false) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savepasaranonline", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page: "PASARAN-SAVE",
                    idpasaran: idcomppasaran,
                    pasaran_hariraya: select_pasaranonline,
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
                    EditData(idcomppasaran,pasaran_tipe)
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_class = "btn btn-primary"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
            }
        } else {
            if(msg_error != ""){
                isModalNotif = true;
            }
        }
    }
    async function deletePasaranOnline(e) {
		const res = await fetch(path_api+"api/deletepasaranonline", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: "Bearer " + token,
			},
			body: JSON.stringify({
				idpasaran: parseInt(idcomppasaran),
				idpasaraonline: parseInt(e),
				page:"PASARAN-SAVE",
			}),
		});
		const json = await res.json();
		if (json.status == 200) {
			EditData(idcomppasaran,pasaran_tipe);
		}else if(json.status == 403){
			alert(json.message)
		}
	}
    async function SaveLimitline() {
        let flag = false;
        msg_error = "";
        if (pasaran_limitline4d_field == "") {
            flag = true;
            msg_error += "The Limitline 4D is required<br>";
        }
        if (pasaran_limitline3d_field == "") {
            flag = true;
            msg_error += "The Limitline 3D is required<br>";
        }
        if (pasaran_limitline3dd_field == "") {
            flag = true;
            msg_error += "The Limitline 3DD is required<br>";
        }
        if (pasaran_limitline2d_field == "") {
            flag = true;
            msg_error += "The Limitline 2D is required<br>";
        }
        if (pasaran_limitline2dd_field == "") {
            flag = true;
            msg_error += "The Limitline 2DD is required<br>";
        }
        if (pasaran_limitline2dt_field == "") {
            flag = true;
            msg_error += "The Limitline 2DT is required<br>";
        }
        if (flag == false) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savepasaranlimitline", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    page: "PASARAN-SAVE",
                    pasaran_limitline4d: parseInt(pasaran_limitline4d_field),
                    pasaran_limitline3d: parseInt(pasaran_limitline3d_field),
                    pasaran_limitline3dd: parseInt(pasaran_limitline3dd_field),
                    pasaran_limitline2d: parseInt(pasaran_limitline2d_field),
                    pasaran_limitline2dd: parseInt(pasaran_limitline2dd_field),
                    pasaran_limitline2dt: parseInt(pasaran_limitline2dt_field),
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
                    EditData(idcomppasaran,pasaran_tipe)
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_class = "btn btn-primary"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
            }
        } else {
            if(msg_error != ""){
                isModalNotif = true;
            }
        }
    }
    const ChangeTabMenu = (e) => {
        switch(e){
            case "menu_pasaranonline":
                tab_pasaranonline = "bg-sky-600 text-white"
                tab_bbfs = ""
                tab_configure = ""
                panel_pasaranonline = true
                panel_bbfs = false
                panel_configure = false
                break;
            case "menu_bbfs":
                tab_pasaranonline = ""
                tab_configure = ""
                tab_bbfs = "bg-sky-600 text-white"
                panel_pasaranonline = false
                panel_configure = false
                panel_bbfs = true
                break;
            case "menu_configure":
                tab_pasaranonline = ""
                tab_bbfs = ""
                tab_configure = "bg-sky-600 text-white"
                panel_pasaranonline = false
                panel_bbfs = false
                panel_configure = true
                break;
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const call_configure = (e) => {
        permainan = ""; 
        switch(e){
            case "4-3-2":
                modalpasaran_width = "max-w-5xl ";
                permainan = "4D/3D/2D"
                isModal_Form_pasaran = true;
                panel_432D = true;
                panel_cbebas = false;
                panel_cmacau = false
                panel_cnaga = false
                panel_cjitu = false
                panel_5050umum = false
                panel_5050special = false
                panel_5050kombinasi = false
                panel_macaukombinasi = false
                panel_dasar = false
                panel_shio = false
                break;
            case "colok_bebas":
                modalpasaran_width = "max-w-3xl ";
                permainan = "Colok Bebas";
                isModal_Form_pasaran = true;
                panel_cbebas = true;
                panel_432D = false;
                panel_cmacau = false
                panel_cnaga = false
                panel_cjitu = false
                panel_5050umum = false
                panel_5050special = false
                panel_5050kombinasi = false
                panel_macaukombinasi = false
                panel_dasar = false
                panel_shio = false
                break;
            case "colok_macau":
                modalpasaran_width = "max-w-3xl ";
                permainan = "Colok Macau";
                isModal_Form_pasaran = true;
                panel_432D = false;
                panel_cbebas = false;
                panel_cmacau = true
                panel_cnaga = false
                panel_cjitu = false
                panel_5050umum = false
                panel_5050special = false
                panel_5050kombinasi = false
                panel_macaukombinasi = false
                panel_dasar = false
                panel_shio = false
                break;
            case "colok_naga":
                modalpasaran_width = "max-w-3xl ";
                permainan = "Colok Macau";
                isModal_Form_pasaran = true;
                panel_432D = false;
                panel_cbebas = false;
                panel_cmacau = false
                panel_cnaga = true
                panel_cjitu = false
                panel_5050umum = false
                panel_5050special = false
                panel_5050kombinasi = false
                panel_macaukombinasi = false
                panel_dasar = false
                panel_shio = false
                break;
            case "colok_jitu":
                modalpasaran_width = "max-w-3xl ";
                permainan = "Colok Jitu";
                isModal_Form_pasaran = true;
                panel_432D = false;
                panel_cbebas = false;
                panel_cmacau = false
                panel_cnaga = false
                panel_cjitu = true;
                panel_5050umum = false
                panel_5050special = false
                panel_5050kombinasi = false
                panel_macaukombinasi = false
                panel_dasar = false
                panel_shio = false
                break;
            case "5050_umum":
                modalpasaran_width = "max-w-3xl ";
                permainan = "5050 Umum";
                isModal_Form_pasaran = true;
                panel_432D = false;
                panel_cbebas = false;
                panel_cmacau = false
                panel_cnaga = false
                panel_cjitu = false;
                panel_5050umum = true
                panel_5050special = false
                panel_5050kombinasi = false
                panel_macaukombinasi = false
                panel_dasar = false
                panel_shio = false
                break;
            case "5050_special":
                modalpasaran_width = "max-w-3xl ";
                permainan = "5050 Special";
                isModal_Form_pasaran = true;
                panel_432D = false;
                panel_cbebas = false;
                panel_cmacau = false
                panel_cnaga = false
                panel_cjitu = false;
                panel_5050umum = false
                panel_5050special = true
                panel_5050kombinasi = false
                panel_macaukombinasi = false
                panel_dasar = false
                panel_shio = false
                break;
            case "5050_kombinasi":
                modalpasaran_width = "max-w-6xl ";
                permainan = "5050 Kombinasi";
                isModal_Form_pasaran = true;
                panel_432D = false;
                panel_cbebas = false;
                panel_cmacau = false
                panel_cnaga = false
                panel_cjitu = false;
                panel_5050umum = false
                panel_5050special = false
                panel_5050kombinasi = true
                panel_macaukombinasi = false
                panel_dasar = false
                panel_shio = false
                break;
            case "macau_kombinasi":
                modalpasaran_width = "max-w-3xl ";
                permainan = "Macau / Kombinasi";
                isModal_Form_pasaran = true;
                panel_432D = false;
                panel_cbebas = false;
                panel_cmacau = false
                panel_cnaga = false
                panel_cjitu = false;
                panel_5050umum = false
                panel_5050special = false
                panel_5050kombinasi = false
                panel_macaukombinasi = true
                panel_dasar = false
                panel_shio = false
                break;
            case "dasar":
                modalpasaran_width = "max-w-3xl ";
                permainan = "Dasar";
                isModal_Form_pasaran = true;
                panel_432D = false;
                panel_cbebas = false;
                panel_cmacau = false
                panel_cnaga = false
                panel_cjitu = false;
                panel_5050umum = false
                panel_5050special = false
                panel_5050kombinasi = false
                panel_macaukombinasi = false
                panel_dasar = true
                panel_shio = false
                break;
            case "shio":
                modalpasaran_width = "max-w-3xl ";
                permainan = "Shio";
                isModal_Form_pasaran = true;
                panel_432D = false;
                panel_cbebas = false;
                panel_cmacau = false
                panel_cnaga = false
                panel_cjitu = false;
                panel_5050umum = false
                panel_5050special = false
                panel_5050kombinasi = false
                panel_macaukombinasi = false
                panel_dasar = false
                panel_shio = true
                break;
        }
    }
    function clearField(){
        tab_pasaranonline = "bg-sky-600 text-white"
        tab_bbfs = ""
        tab_configure = ""
        panel_pasaranonline = true
        panel_bbfs = false
        panel_configure = false
        pasaran_display_field = 0;
        pasaran_status_field = "";
        pasaran_create_field = "";
        pasaran_createdate_field = "";
        pasaran_update_field = "";
        pasaran_updatedate_field = "";
        $form.form_pasaran_situs_field = "";
        $form.form_pasaran_diundi_field = "";
        $form.form_pasaran_tutup_field = "";
        $form.form_pasaran_jadwal_field = "";
        $form.form_pasaran_open_field = "";
        pasaran_limitline4d_field = 0;
        pasaran_limitline3d_field = 0;
        pasaran_limitline3dd_field = 0;
        pasaran_limitline2d_field = 0;
        pasaran_limitline2dd_field = 0;
        pasaran_limitline2dt_field = 0;
        pasaran_bbfs_field = 0;

        pasaran_minbet_cbebas_field = 0;
        pasaran_maxbet_cbebas_field = 0;
        pasaran_limitotal_cbebas_field = 0;
        pasaran_limitglobal_cbebas_field = 0;
        pasaran_win_cbebas_field = 0;
        pasaran_disc_cbebas_field = 0;
        pasaran_minbet_cmacau_field = 0;
        pasaran_maxbet_cmacau_field = 0;
        pasaran_limitotal_cmacau_field = 0;
        pasaran_limitglobal_cmacau_field = 0;
        pasaran_win2_cmacau_field = 0;
        pasaran_win3_cmacau_field = 0;
        pasaran_win4_cmacau_field = 0;
        pasaran_disc_cmacau_field = 0;
        pasaran_minbet_cnaga_field = 0;
        pasaran_maxbet_cnaga_field = 0;
        pasaran_win3_cnaga_field = 0;
        pasaran_win4_cnaga_field = 0;
        pasaran_disc_cnaga_field = 0;
        pasaran_limitglobal_cnaga_field = 0;
        pasaran_limittotal_cnaga_field = 0;
        pasaran_minbet_cjitu_field = 0;
        pasaran_maxbet_cjitu_field = 0;
        pasaran_winas_cjitu_field = 0;
        pasaran_winkop_cjitu_field = 0;
        pasaran_winkepala_cjitu_field = 0;
        pasaran_winekor_cjitu_field = 0;
        pasaran_desc_cjitu_field = 0;
        pasaran_limitglobal_cjitu_field = 0;
        pasaran_limittotal_cjitu_field = 0;
        pasaran_minbet_5050umum_field = 0;
        pasaran_maxbet_5050umum_field = 0;
        pasaran_keibesar_5050umum_field = 0;
        pasaran_keikecil_5050umum_field = 0;
        pasaran_keigenap_5050umum_field = 0;
        pasaran_keiganjil_5050umum_field = 0;
        pasaran_keitengah_5050umum_field = 0;
        pasaran_keitepi_5050umum_field = 0;
        pasaran_discbesar_5050umum_field = 0;
        pasaran_disckecil_5050umum_field = 0;
        pasaran_discgenap_5050umum_field = 0;
        pasaran_discganjil_5050umum_field = 0;
        pasaran_disctengah_5050umum_field = 0;
        pasaran_disctepi_5050umum_field = 0;
        pasaran_limitglobal_5050umum_field = 0;
        pasaran_limittotal_5050umum_field = 0;
        pasaran_minbet_5050special_field = 0;
        pasaran_maxbet_5050special_field = 0;
        pasaran_keiasganjil_5050special_field = 0;
        pasaran_keiasgenap_5050special_field = 0;
        pasaran_keiasbesar_5050special_field = 0;
        pasaran_keiaskecil_5050special_field = 0;
        pasaran_keikopganjil_5050special_field = 0;
        pasaran_keikopgenap_5050special_field = 0;
        pasaran_keikopbesar_5050special_field = 0;
        pasaran_keikopkecil_5050special_field = 0;
        pasaran_keikepalaganjil_5050special_field = 0;
        pasaran_keikepalagenap_5050special_field = 0;
        pasaran_keikepalabesar_5050special_field = 0;
        pasaran_keikepalakecil_5050special_field = 0;
        pasaran_keiekorganjil_5050special_field = 0;
        pasaran_keiekorgenap_5050special_field = 0;
        pasaran_keiekorbesar_5050special_field = 0;
        pasaran_keiekorkecil_5050special_field = 0;
        pasaran_discasganjil_5050special_field = 0;
        pasaran_discasgenap_5050special_field = 0;
        pasaran_discasbesar_5050special_field = 0;
        pasaran_discaskecil_5050special_field = 0;
        pasaran_disckopganjil_5050special_field = 0;
        pasaran_disckopgenap_5050special_field = 0;
        pasaran_disckopbesar_5050special_field = 0;
        pasaran_disckopkecil_5050special_field = 0;
        pasaran_disckepalaganjil_5050special_field = 0;
        pasaran_disckepalagenap_5050special_field = 0;
        pasaran_disckepalabesar_5050special_field = 0;
        pasaran_disckepalakecil_5050special_field = 0;
        pasaran_discekorganjil_5050special_field = 0;
        pasaran_discekorgenap_5050special_field = 0;
        pasaran_discekorbesar_5050special_field = 0;
        pasaran_discekorkecil_5050special_field = 0;
        pasaran_limitglobal_5050special_field = 0;
        pasaran_limittotal_5050special_field = 0;
        pasaran_minbet_5050kombinasi_field = 0;
        pasaran_maxbet_5050kombinasi_field = 0;
        pasaran_belakangkeimono_5050kombinasi_field = 0;
        pasaran_belakangkeistereo_5050kombinasi_field = 0;
        pasaran_belakangkeikembang_5050kombinasi_field = 0;
        pasaran_belakangkeikempis_5050kombinasi_field = 0;
        pasaran_belakangkeikembar_5050kombinasi_field = 0;
        pasaran_tengahkeimono_5050kombinasi_field = 0;
        pasaran_tengahkeistereo_5050kombinasi_field = 0;
        pasaran_tengahkeikembang_5050kombinasi_field = 0;
        pasaran_tengahkeikempis_5050kombinasi_field = 0;
        pasaran_tengahkeikembar_5050kombinasi_field = 0;
        pasaran_depankeimono_5050kombinasi_field = 0;
        pasaran_depankeistereo_5050kombinasi_field = 0;
        pasaran_depankeikembang_5050kombinasi_field = 0;
        pasaran_depankeikempis_5050kombinasi_field = 0;
        pasaran_depankeikembar_5050kombinasi_field = 0;
        pasaran_belakangdiscmono_5050kombinasi_field = 0;
        pasaran_belakangdiscstereo_5050kombinasi_field = 0;
        pasaran_belakangdisckembang_5050kombinasi_field = 0;
        pasaran_belakangdisckempis_5050kombinasi_field = 0;
        pasaran_belakangdisckembar_5050kombinasi_field = 0;
        pasaran_tengahdiscmono_5050kombinasi_field = 0;
        pasaran_tengahdiscstereo_5050kombinasi_field = 0;
        pasaran_tengahdisckembang_5050kombinasi_field = 0;
        pasaran_tengahdisckempis_5050kombinasi_field = 0;
        pasaran_tengahdisckembar_5050kombinasi_field = 0;
        pasaran_depandiscmono_5050kombinasi_field = 0;
        pasaran_depandiscstereo_5050kombinasi_field = 0;
        pasaran_depandisckembang_5050kombinasi_field = 0;
        pasaran_depandisckempis_5050kombinasi_field = 0;
        pasaran_depandisckembar_5050kombinasi_field = 0;
        pasaran_limitglobal_5050kombinasi_field = 0;
        pasaran_limittotal_5050kombinasi_field = 0;
        pasaran_minbet_kombinasi_field = 0;
        pasaran_maxbet_kombinasi_field = 0;
        pasaran_win_kombinasi_field = 0;
        pasaran_disc_kombinasi_field = 0;
        pasaran_limitglobal_kombinasi_field = 0;
        pasaran_limittotal_kombinasi_field = 0;
        pasaran_minbet_dasar_field = 0;
        pasaran_maxbet_dasar_field = 0;
        pasaran_keibesar_dasar_field = 0;
        pasaran_keikecil_dasar_field = 0;
        pasaran_keigenap_dasar_field = 0;
        pasaran_keiganjil_dasar_field = 0;
        pasaran_discbesar_dasar_field = 0;
        pasaran_disckecil_dasar_field = 0;
        pasaran_discgenap_dasar_field = 0;
        pasaran_discganjil_dasar_field = 0;
        pasaran_limitglobal_dasar_field = 0;
        pasaran_limittotal_dasar_field = 0;
        pasaran_minbet_shio_field = 0;
        pasaran_maxbet_shio_field = 0;
        pasaran_win_shio_field = 0;
        pasaran_disc_shio_field = 0;
        pasaran_shioyear_shio_field = "";
        pasaran_limitglobal_shio_field = 0;
        pasaran_limittotal_shio_field = 0;
    }
    const LoadingRunning = () => {
        loader_class = "inline-block"
        loader_msg = "Sending..."
    };
    const LoadingRunningFinish = (e) => {
        loader_msg =e.detail.temp_msg
        setTimeout(function () {
            loader_class = "hidden";
        }, 1000);
    };
    const call_notif = (e) => {
        msg_error = e.detail.temp_msg
        isModalNotif = true;
    };        
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
    on:handleRefresh={RefreshHalaman}
    panel_button_new={false}
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
            type="text" placeholder="Search by Pasaran" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 focus:ring-0 focus:outline-none">
    </slot:template>
    <slot:template slot="panel_body">
        <table class="table table-compact w-full">
            <thead class="sticky top-0">
                <tr>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center"></th>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center" colspan=2>STATUS</th>
                    <th width="*" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">PASARAN</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">HARI DIUNDI</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">TUTUP</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">JADWAL</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">OPEN</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">DISPLAY</th>
                </tr>
            </thead>
            {#if filterHome != ""}
                <tbody>
                    {#each filterHome as rec}
                    <tr>
                        <td on:click={() => {
                            EditData(rec.home_id,rec.home_tipe);
                            }} class="text-center text-xs cursor-pointer">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                            </svg>
                        </td>
                        <td class="text-xs lg:text-sm align-top text-center">
                            <span class="{rec.home_status_class} text-center rounded-md p-1 px-2 shadow-lg ">{rec.home_status}</span>
                        </td>
                        <td class="text-xs lg:text-sm align-top text-center">
                            <span class="{rec.home_active_class} text-center rounded-md p-1 px-2 shadow-lg ">{rec.home_active}</span>
                        </td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_nama}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_diundi}</td>
                        <td class="text-xs lg:text-sm align-top text-center">{rec.home_jamtutup}</td>
                        <td class="text-xs lg:text-sm align-top text-center">{rec.home_jamjadwal}</td>
                        <td class="text-xs lg:text-sm align-top text-center">{rec.home_jamopen}</td>
                        <td class="text-xs lg:text-sm align-top text-right">{rec.home_display}</td>
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
    modal_popup_class="select-none w-11/12 {modal_width} overflow-hidden">
    <slot:template slot="modalpopup_body">
        <div class="flex justify-between gap-2">
            <div class="w-full">
                <div class="grid grid-cols-2 gap-1">
                    <div class="mb-5">
                        <Input_custom
                            input_onchange="{handleChange}"
                            input_autofocus={false}
                            input_required={false}
                            input_enabled={false}
                            input_tipe="text"
                            bind:value={pasaran_name_field}
                            input_id="pasaran_name_field"
                            input_placeholder="Pasaran"/>
                    </div>
                    <div class="mb-5">
                        <Input_custom
                            input_onchange="{handleChange}"
                            input_autofocus={false}
                            input_required={true}
                            input_tipe="time"
                            input_invalid={$errors.form_pasaran_tutup_field.length > 0}
                            bind:value={$form.form_pasaran_tutup_field}
                            input_id="form_pasaran_tutup_field"
                            input_placeholder="Tutup"/>
                        {#if $errors.form_pasaran_tutup_field}
                            <small class="text-pink-600 text-[11px]">{$errors.form_pasaran_tutup_field}</small>
                        {/if}
                    </div>
                    <div class="mb-5">
                        <Input_custom
                            input_onchange="{handleChange}"
                            input_autofocus={false}
                            input_required={true}
                            input_tipe="text"
                            input_invalid={$errors.form_pasaran_situs_field.length > 0}
                            bind:value={$form.form_pasaran_situs_field}
                            input_id="form_pasaran_situs_field"
                            input_placeholder="Situs"/>
                        {#if $errors.form_pasaran_situs_field}
                            <small class="text-pink-600 text-[11px]">{$errors.form_pasaran_situs_field}</small>
                        {/if}
                    </div>
                    <div class="mb-5">
                        <Input_custom
                            input_onchange="{handleChange}"
                            input_autofocus={false}
                            input_required={true}
                            input_tipe="time"
                            input_invalid={$errors.form_pasaran_jadwal_field.length > 0}
                            bind:value={$form.form_pasaran_jadwal_field}
                            input_id="form_pasaran_jadwal_field"
                            input_placeholder="Jadwal"/>
                        {#if $errors.form_pasaran_jadwal_field}
                            <small class="text-pink-600 text-[11px]">{$errors.form_pasaran_jadwal_field}</small>
                        {/if}
                    </div>
                    <div class="mb-5">
                        <Input_custom
                            input_onchange="{handleChange}"
                            input_autofocus={false}
                            input_required={true}
                            input_tipe="text"
                            input_invalid={$errors.form_pasaran_diundi_field.length > 0}
                            bind:value={$form.form_pasaran_diundi_field}
                            input_id="form_pasaran_diundi_field"
                            input_placeholder="Diundi"/>
                        {#if $errors.form_pasaran_diundi_field}
                            <small class="text-pink-600 text-[11px]">{$errors.form_pasaran_tutup_field}</small>
                        {/if}
                    </div>
                    <div class="mb-5">
                        <Input_custom
                            input_onchange="{handleChange}"
                            input_autofocus={false}
                            input_required={true}
                            input_tipe="time"
                            input_invalid={$errors.form_pasaran_open_field.length > 0}
                            bind:value={$form.form_pasaran_open_field}
                            input_id="form_pasaran_open_field"
                            input_placeholder="Buka"/>
                        {#if $errors.form_pasaran_open_field}
                            <small class="text-pink-600 text-[11px]">{$errors.form_pasaran_open_field}</small>
                        {/if}
                    </div>
                    <div class="mb-5 text-[11px]">
                        Create : {pasaran_create_field}, {pasaran_createdate_field}
                        {#if pasaran_update_field != "" }
                            <br>
                            Update : {pasaran_update_field}, {pasaran_updatedate_field}
                        {/if}
                    </div>
                    <div class="mb-5">
                        <Input_custom
                            input_autofocus={false}
                            input_required={true}
                            input_tipe="number_nolabel"
                            bind:value={pasaran_display_field}
                            input_id="pasaran_display_field"
                            input_placeholder="Display"/>
                    </div>
                    <div class="mb-5">
                        
                    </div>
                    <div class="mb-5">
                        <select
                            bind:value={pasaran_status_field}
                            class="w-full rounded px-3  border border-gray-300 focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none">
                            <option value="Y">ACTIVE</option>
                            <option value="N">DEACTIVE</option>
                        </select>
                    </div>
                    {#if pasaran_tipe != "WAJIB"}
                    <div class="col-span-2">
                        <button on:click={() => {
                            handleSubmit();
                        }} class="{buttonLoading_class} btn-block">Submit</button>
                    </div>
                    {/if}
                </div>
            </div>
            <div class="w-full p-2 -mt-5">
                <ul class="flex justify-center items-center gap-2">
                    <li on:click={() => {
                            ChangeTabMenu("menu_pasaranonline");
                        }}
                        class="items-center {tab_pasaranonline}  px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">Pasaran Online</li>
                    <li on:click={() => {
                            ChangeTabMenu("menu_bbfs");
                        }}
                        class="items-center {tab_bbfs} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">Bolak Balik</li>
                    <li on:click={() => {
                            ChangeTabMenu("menu_configure");
                        }}
                        class="items-center {tab_configure} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">Konfigurasi</li>
                </ul>
                {#if panel_pasaranonline}
                    <div class="grid grid-cols-1 gap-2">
                        <h2 class="text-lg font-bold">Setting Pasaran Online</h2>
                        {#if pasaran_tipe != "WAJIB"}
                        <div class="form-control">
                            <div class="input-group">
                            <select bind:value={select_pasaranonline} class="select select-bordered w-[80%]">
                                <option disabled selected value="">--Pilih Hari--</option>
                                <option value="senin">Senin</option>
                                <option value="selasa">Selasa</option>
                                <option value="rabu">Rabu</option>
                                <option value="kamis">Kamis</option>
                                <option value="jumat">Jumat</option>
                                <option value="sabtu">Sabtu</option>
                                <option value="minggu">Minggu</option>
                            </select>
                            <button on:click={() => {
                                    savePasaranOnline();
                                }} class="{buttonLoading2_class} btn-primary">Save</button>
                            </div>
                        </div>
                        {/if}
                        <div>
                            <table class="table table-compact w-full">
                                <tbody>
                                    {#if listPasaranOnline != ""}
                                        {#each listPasaranOnline as rec}
                                            <tr>
                                                {#if pasaran_tipe != "WAJIB"}
                                                <td on:click={() => {
                                                        deletePasaranOnline(rec.idpasaranonline);
                                                    }} width="2%" class="text-center align-top cursor-pointer">
                                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                                        <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                                                    </svg>
                                                </td>
                                                {/if}
                                                <td width="*" class="text-left align-top capitalize">{rec.haripasaran}</td>
                                            </tr>
                                        {/each}
                                    {:else}
                                        <tr>
                                            <td colspan="2" class="text-left text-xs font-semibold">
                                                No Record
                                            </td>
                                        </tr>
                                    {/if}
                                </tbody>
                            </table>
                        </div>
                    </div>
                {/if}
                {#if panel_bbfs}
                    <div class="grid grid-cols-3 gap-2">
                        <h2 class="text-lg font-bold col-span-3">Setting Limit Line + Bolak Balik</h2>
                        <Input_custom
                            input_enabled={false}
                            input_tipe="number_nolabel"
                            bind:value={pasaran_bbfs_field}
                            input_id="pasaran_bbfs_field"
                            input_placeholder="BBFS"/>
                        <Input_custom
                            input_enabled={true}
                            input_tipe="number"
                            bind:value={pasaran_limitline4d_field}
                            input_id="pasaran_limitline4d_field"
                            input_placeholder="LimitLine 4D"/>
                        <Input_custom
                            input_enabled={true}
                            input_tipe="number"
                            bind:value={pasaran_limitline2d_field}
                            input_id="pasaran_limitline2d_field"
                            input_placeholder="LimitLine 2D"/>
                        <div></div>
                        <Input_custom
                            input_enabled={true}
                            input_tipe="number"
                            bind:value={pasaran_limitline3d_field}
                            input_id="pasaran_limitline3d_field"
                            input_placeholder="LimitLine 3D"/>
                        <Input_custom
                            input_enabled={true}
                            input_tipe="number"
                            bind:value={pasaran_limitline2dd_field}
                            input_id="pasaran_limitline2dd_field"
                            input_placeholder="LimitLine 2DD"/>
                        <div></div>
                        <Input_custom
                            input_enabled={true}
                            input_tipe="number"
                            bind:value={pasaran_limitline3dd_field}
                            input_id="pasaran_limitline3dd_field"
                            input_placeholder="LimitLine 3DD"/>
                        <Input_custom
                            input_enabled={true}
                            input_tipe="number"
                            bind:value={pasaran_limitline2dt_field}
                            input_id="pasaran_limitline2dt_field"
                            input_placeholder="LimitLine 2DT"/>
                        {#if pasaran_tipe != "WAJIB"}
                        <div class="col-span-3">
                            <button on:click={() => {
                                SaveLimitline();
                            }} class="{buttonLoading_class} btn-block">Submit</button>
                        </div>
                        {/if}
                    </div>
                {/if}
                {#if panel_configure}
                    <div class="grid grid-cols-3 gap-2">
                        <h2 class="text-lg font-bold col-span-3">Configure</h2>
                        <button on:click={() => {
                            call_configure("4-3-2");
                        }} class="btn btn-warning">4D/3D/2D</button>
                        <button on:click={() => {
                            call_configure("colok_bebas");
                        }} class="btn btn-warning">COLOK BEBAS</button>
                        <button on:click={() => {
                            call_configure("colok_macau");
                        }} class="btn btn-warning">COLOK MACAU</button>
                        <button on:click={() => {
                            call_configure("colok_naga");
                        }} class="btn btn-warning">COLOK NAGA</button>
                        <button on:click={() => {
                            call_configure("colok_jitu");
                        }} class="btn btn-warning">COLOK JITU</button>
                        <button on:click={() => {
                            call_configure("5050_umum");
                        }} class="btn btn-warning">5050 UMUM</button>
                        <button on:click={() => {
                            call_configure("5050_special");
                        }} class="btn btn-warning">5050 SPECIAL</button>
                        <button on:click={() => {
                            call_configure("5050_kombinasi");
                        }} class="btn btn-warning">5050 KOMBINASI</button>
                        <button on:click={() => {
                            call_configure("macau_kombinasi");
                        }} class="btn btn-warning">MACAU / KOMBINASI</button>
                        <button on:click={() => {
                            call_configure("dasar");
                        }} class="btn btn-warning">DASAR</button>
                        <button on:click={() => {
                            call_configure("shio");
                        }} class="btn btn-warning">SHIO</button>
                    </div>
                {/if}
            </div>
        </div>
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-formpasaran" class="modal-toggle" bind:checked={isModal_Form_pasaran}>
<Modal_popup
    modal_popup_id="my-modal-formpasaran"
    modal_popup_title="{permainan}"
    modal_popup_class="select-none w-11/12 {modalpasaran_width} scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100">
    <slot:template slot="modalpopup_body">
        <div class="flex flex-col">
            {#if panel_432D}
                <Panel_432D
                    on:handleLoadingRunning={LoadingRunning}
                    on:handleLoadingRunningFinish={LoadingRunningFinish}
                    on:handleCallNotif={call_notif}
                    {pasaran_tipe}
                    {path_api}
                    {token}
                    {idcomppasaran}
                    {pasaran_idpasarantogel_field}
                    {pasaran_minbet_432d_field}
                    {pasaran_maxbet4d_432d_field}
                    {pasaran_maxbet3d_432d_field}
                    {pasaran_maxbet3dd_432d_field}
                    {pasaran_maxbet2d_432d_field}
                    {pasaran_maxbet2dd_432d_field}
                    {pasaran_maxbet2dt_432d_field}
                    {pasaran_limitotal4d_432d_field}
                    {pasaran_limitotal3d_432d_field}
                    {pasaran_limitotal3dd_432d_field}
                    {pasaran_limitotal2d_432d_field}
                    {pasaran_limitotal2dd_432d_field}
                    {pasaran_limitotal2dt_432d_field}
                    {pasaran_limitglobal4d_432d_field}
                    {pasaran_limitglobal3d_432d_field}
                    {pasaran_limitglobal3dd_432d_field}
                    {pasaran_limitglobal2d_432d_field}
                    {pasaran_limitglobal2dd_432d_field}
                    {pasaran_limitglobal2dt_432d_field}
                    {pasaran_disc4d_432d_field}
                    {pasaran_disc3d_432d_field}
                    {pasaran_disc3dd_432d_field}
                    {pasaran_disc2d_432d_field}
                    {pasaran_disc2dd_432d_field}
                    {pasaran_disc2dt_432d_field}
                    {pasaran_win4d_432d_field}
                    {pasaran_win3d_432d_field}
                    {pasaran_win3dd_432d_field}
                    {pasaran_win2d_432d_field}
                    {pasaran_win2dd_432d_field}
                    {pasaran_win2dt_432d_field}
                    {pasaran_win4dnodisc_432d_field}
                    {pasaran_win3dnodisc_432d_field}
                    {pasaran_win3ddnodisc_432d_field}
                    {pasaran_win2dnodisc_432d_field}
                    {pasaran_win2ddnodisc_432d_field}
                    {pasaran_win2dtnodisc_432d_field}
                    {pasaran_win4d_bb_kena_432d_field}
                    {pasaran_win3d_bb_kena_432d_field}
                    {pasaran_win3dd_bb_kena_432d_field}
                    {pasaran_win2d_bb_kena_432d_field}
                    {pasaran_win2dd_bb_kena_432d_field}
                    {pasaran_win2dt_bb_kena_432d_field}
                    {pasaran_win4d_bb_432d_field}
                    {pasaran_win3d_bb_432d_field}
                    {pasaran_win3dd_bb_432d_field}
                    {pasaran_win2d_bb_432d_field}
                    {pasaran_win2dd_bb_432d_field}
                    {pasaran_win2dt_bb_432d_field}
                 />
            {/if}
            {#if panel_cbebas}
                <Panel_colokbebas
                    on:handleLoadingRunning={LoadingRunning}
                    on:handleLoadingRunningFinish={LoadingRunningFinish}
                    on:handleCallNotif={call_notif}
                    {pasaran_tipe}
                    {path_api}
                    {token}
                    {idcomppasaran}
                    {pasaran_idpasarantogel_field}
                    {pasaran_minbet_cbebas_field}
                    {pasaran_maxbet_cbebas_field}
                    {pasaran_limitotal_cbebas_field}
                    {pasaran_limitglobal_cbebas_field}
                    {pasaran_win_cbebas_field}
                    {pasaran_disc_cbebas_field} />
            {/if}
            {#if panel_cmacau}
                <Panel_colokmacau
                    on:handleLoadingRunning={LoadingRunning}
                    on:handleLoadingRunningFinish={LoadingRunningFinish}
                    on:handleCallNotif={call_notif}
                    {pasaran_tipe}
                    {path_api}
                    {token}
                    {idcomppasaran}
                    {pasaran_idpasarantogel_field}
                    {pasaran_minbet_cmacau_field}
                    {pasaran_maxbet_cmacau_field}
                    {pasaran_limitotal_cmacau_field}
                    {pasaran_limitglobal_cmacau_field}
                    {pasaran_win2_cmacau_field}
                    {pasaran_win3_cmacau_field}
                    {pasaran_win4_cmacau_field}
                    {pasaran_disc_cmacau_field} />
            {/if}
            {#if panel_cnaga}
                <Panel_coloknaga
                    on:handleLoadingRunning={LoadingRunning}
                    on:handleLoadingRunningFinish={LoadingRunningFinish}
                    on:handleCallNotif={call_notif}
                    {pasaran_tipe}
                    {path_api}
                    {token}
                    {idcomppasaran}
                    {pasaran_idpasarantogel_field}
                    {pasaran_minbet_cnaga_field}
                    {pasaran_maxbet_cnaga_field}
                    {pasaran_win3_cnaga_field}
                    {pasaran_win4_cnaga_field}
                    {pasaran_disc_cnaga_field}
                    {pasaran_limitglobal_cnaga_field}
                    {pasaran_limittotal_cnaga_field} />
            {/if}
            {#if panel_cjitu}
                <Panel_colokjitu
                    on:handleLoadingRunning={LoadingRunning}
                    on:handleLoadingRunningFinish={LoadingRunningFinish}
                    on:handleCallNotif={call_notif}
                    {pasaran_tipe}
                    {path_api}
                    {token}
                    {idcomppasaran}
                    {pasaran_idpasarantogel_field}
                    {pasaran_minbet_cjitu_field}
                    {pasaran_maxbet_cjitu_field}
                    {pasaran_winas_cjitu_field}
                    {pasaran_winkop_cjitu_field}
                    {pasaran_winkepala_cjitu_field}
                    {pasaran_winekor_cjitu_field}
                    {pasaran_desc_cjitu_field}
                    {pasaran_limitglobal_cjitu_field}
                    {pasaran_limittotal_cjitu_field} />
            {/if}
            {#if panel_5050umum}
                <Panel_5050umum
                    on:handleLoadingRunning={LoadingRunning}
                    on:handleLoadingRunningFinish={LoadingRunningFinish}
                    on:handleCallNotif={call_notif}
                    {pasaran_tipe}
                    {path_api}
                    {token}
                    {idcomppasaran}
                    {pasaran_idpasarantogel_field}
                    {pasaran_minbet_5050umum_field}
                    {pasaran_maxbet_5050umum_field}
                    {pasaran_keibesar_5050umum_field}
                    {pasaran_keikecil_5050umum_field}
                    {pasaran_keigenap_5050umum_field}
                    {pasaran_keiganjil_5050umum_field}
                    {pasaran_keitengah_5050umum_field}
                    {pasaran_keitepi_5050umum_field}
                    {pasaran_discbesar_5050umum_field}
                    {pasaran_disckecil_5050umum_field}
                    {pasaran_discgenap_5050umum_field}
                    {pasaran_discganjil_5050umum_field}
                    {pasaran_disctengah_5050umum_field}
                    {pasaran_disctepi_5050umum_field}
                    {pasaran_limitglobal_5050umum_field}
                    {pasaran_limittotal_5050umum_field} />
            {/if}
            {#if panel_5050special}
                <Panel_5050special
                    on:handleLoadingRunning={LoadingRunning}
                    on:handleLoadingRunningFinish={LoadingRunningFinish}
                    on:handleCallNotif={call_notif}
                    {pasaran_tipe}
                    {path_api}
                    {token}
                    {idcomppasaran}
                    {pasaran_idpasarantogel_field}
                    {pasaran_minbet_5050special_field}
                    {pasaran_maxbet_5050special_field}
                    {pasaran_keiasganjil_5050special_field}
                    {pasaran_keiasgenap_5050special_field}
                    {pasaran_keiasbesar_5050special_field}
                    {pasaran_keiaskecil_5050special_field}
                    {pasaran_keikopganjil_5050special_field}
                    {pasaran_keikopgenap_5050special_field}
                    {pasaran_keikopbesar_5050special_field}
                    {pasaran_keikopkecil_5050special_field}
                    {pasaran_keikepalaganjil_5050special_field}
                    {pasaran_keikepalagenap_5050special_field}
                    {pasaran_keikepalabesar_5050special_field}
                    {pasaran_keikepalakecil_5050special_field}
                    {pasaran_keiekorganjil_5050special_field}
                    {pasaran_keiekorgenap_5050special_field}
                    {pasaran_keiekorbesar_5050special_field}
                    {pasaran_keiekorkecil_5050special_field}
                    {pasaran_discasganjil_5050special_field}
                    {pasaran_discasgenap_5050special_field}
                    {pasaran_discasbesar_5050special_field}
                    {pasaran_discaskecil_5050special_field}
                    {pasaran_disckopganjil_5050special_field}
                    {pasaran_disckopgenap_5050special_field}
                    {pasaran_disckopbesar_5050special_field}
                    {pasaran_disckopkecil_5050special_field}
                    {pasaran_disckepalaganjil_5050special_field}
                    {pasaran_disckepalagenap_5050special_field}
                    {pasaran_disckepalabesar_5050special_field}
                    {pasaran_disckepalakecil_5050special_field}
                    {pasaran_discekorganjil_5050special_field}
                    {pasaran_discekorgenap_5050special_field}
                    {pasaran_discekorbesar_5050special_field}
                    {pasaran_discekorkecil_5050special_field}
                    {pasaran_limitglobal_5050special_field}
                    {pasaran_limittotal_5050special_field} />
            {/if}
            {#if panel_5050kombinasi}
                <Panel_5050kombinasi
                    on:handleLoadingRunning={LoadingRunning}
                    on:handleLoadingRunningFinish={LoadingRunningFinish}
                    on:handleCallNotif={call_notif}
                    {pasaran_tipe}
                    {path_api}
                    {token}
                    {idcomppasaran}
                    {pasaran_idpasarantogel_field}
                    {pasaran_minbet_5050kombinasi_field}
                    {pasaran_maxbet_5050kombinasi_field}
                    {pasaran_belakangkeimono_5050kombinasi_field}
                    {pasaran_belakangkeistereo_5050kombinasi_field}
                    {pasaran_belakangkeikembang_5050kombinasi_field}
                    {pasaran_belakangkeikempis_5050kombinasi_field}
                    {pasaran_belakangkeikembar_5050kombinasi_field}
                    {pasaran_tengahkeimono_5050kombinasi_field}
                    {pasaran_tengahkeistereo_5050kombinasi_field}
                    {pasaran_tengahkeikembang_5050kombinasi_field}
                    {pasaran_tengahkeikempis_5050kombinasi_field}
                    {pasaran_tengahkeikembar_5050kombinasi_field}
                    {pasaran_depankeimono_5050kombinasi_field}
                    {pasaran_depankeistereo_5050kombinasi_field}
                    {pasaran_depankeikembang_5050kombinasi_field}
                    {pasaran_depankeikempis_5050kombinasi_field}
                    {pasaran_depankeikembar_5050kombinasi_field}
                    {pasaran_belakangdiscmono_5050kombinasi_field}
                    {pasaran_belakangdiscstereo_5050kombinasi_field}
                    {pasaran_belakangdisckembang_5050kombinasi_field}
                    {pasaran_belakangdisckempis_5050kombinasi_field}
                    {pasaran_belakangdisckembar_5050kombinasi_field}
                    {pasaran_tengahdiscmono_5050kombinasi_field}
                    {pasaran_tengahdiscstereo_5050kombinasi_field}
                    {pasaran_tengahdisckembang_5050kombinasi_field}
                    {pasaran_tengahdisckempis_5050kombinasi_field}
                    {pasaran_tengahdisckembar_5050kombinasi_field}
                    {pasaran_depandiscmono_5050kombinasi_field}
                    {pasaran_depandiscstereo_5050kombinasi_field}
                    {pasaran_depandisckembang_5050kombinasi_field}
                    {pasaran_depandisckempis_5050kombinasi_field}
                    {pasaran_depandisckembar_5050kombinasi_field}
                    {pasaran_limitglobal_5050kombinasi_field}
                    {pasaran_limittotal_5050kombinasi_field} />
            {/if}
            {#if panel_macaukombinasi}
                <Panel_macaukombinasi
                    on:handleLoadingRunning={LoadingRunning}
                    on:handleLoadingRunningFinish={LoadingRunningFinish}
                    on:handleCallNotif={call_notif}
                    {pasaran_tipe}
                    {path_api}
                    {token}
                    {idcomppasaran}
                    {pasaran_idpasarantogel_field}
                    {pasaran_minbet_kombinasi_field}
                    {pasaran_maxbet_kombinasi_field}
                    {pasaran_win_kombinasi_field}
                    {pasaran_disc_kombinasi_field}
                    {pasaran_limitglobal_kombinasi_field}
                    {pasaran_limittotal_kombinasi_field} />
            {/if}
            {#if panel_dasar}
                <Panel_dasar
                    on:handleLoadingRunning={LoadingRunning}
                    on:handleLoadingRunningFinish={LoadingRunningFinish}
                    on:handleCallNotif={call_notif}
                    {pasaran_tipe}
                    {path_api}
                    {token}
                    {idcomppasaran}
                    {pasaran_idpasarantogel_field}
                    {pasaran_minbet_dasar_field}
                    {pasaran_maxbet_dasar_field}
                    {pasaran_keibesar_dasar_field}
                    {pasaran_keikecil_dasar_field}
                    {pasaran_keigenap_dasar_field}
                    {pasaran_keiganjil_dasar_field}
                    {pasaran_discbesar_dasar_field}
                    {pasaran_disckecil_dasar_field}
                    {pasaran_discgenap_dasar_field}
                    {pasaran_discganjil_dasar_field}
                    {pasaran_limitglobal_dasar_field}
                    {pasaran_limittotal_dasar_field} />
            {/if}
            {#if panel_shio}
                <Panel_shio
                    on:handleLoadingRunning={LoadingRunning}
                    on:handleLoadingRunningFinish={LoadingRunningFinish}
                    on:handleCallNotif={call_notif}
                    {pasaran_tipe}
                    {path_api}
                    {token}
                    {idcomppasaran}
                    {pasaran_idpasarantogel_field}
                    {pasaran_minbet_shio_field}
                    {pasaran_maxbet_shio_field}
                    {pasaran_win_shio_field}
                    {pasaran_disc_shio_field}
                    {pasaran_shioyear_shio_field}
                    {pasaran_limitglobal_shio_field}
                    {pasaran_limittotal_shio_field} />
            {/if}
        </div>
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />
