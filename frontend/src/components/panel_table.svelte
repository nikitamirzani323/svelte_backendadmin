<script>
    export let panel_height = "550px";
    export let panel_class = "";
</script>
<div class="hidden {panel_class} sm:inline w-full scrollbar-thin scrollbar-thumb-sky-200 h-[{panel_height}] ">
    <slot name="paneltable_body" />
</div>