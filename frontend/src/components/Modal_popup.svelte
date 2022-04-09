<script>
    export let modal_popup_id = ""
    export let modal_popup_title = ""
    export let modal_popup_class = "select-none max-w-full lg:max-w-xl overflow-hidden"
</script>

<div class="modal" >
    <div class="modal-box relative    {modal_popup_class} rounded-none lg:rounded-lg p-2  ">
        <div class="flex flex-col items-stretch">
            <div class="h-8">
                <label for="{modal_popup_id}" class="btn btn-xs lg:btn-sm btn-circle absolute right-2 top-2">✕</label>
                <h3 class="text-xs lg:text-sm font-bold mt-1">{modal_popup_title}</h3>
            </div>
            <slot name="modalpopup_body" />
        </div>
    </div>
</div>