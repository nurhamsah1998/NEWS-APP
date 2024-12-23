<template>
   <button class="bg-cyan-500 hover:bg-cyan-600 active:bg-cyan-400 outline-cyan-600 hover:outline-cyan-500 active:scale-[0.9] leading-3 py-3 transition duration-200 px-5 font-semibold text-white rounded-md">{{ textContent }}</button> 
</template>
<script setup lang="ts">
defineProps<{textContent:string}>()
</script>