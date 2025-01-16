<script setup lang="ts">
import NotFoundImage from '@/assets/page-not-found.svg'
</script>

<template>
<div class=" w-full h-dvh flex flex-col gap-5 justify-center items-center">
    <img class=" max-w-36" v-bind:src="NotFoundImage" />
   <div class="text-center">
    <p class=" text-xl font-semibold text-slate-600">Opps... Page Not Found</p>
    <p class=" text-sm text-slate-600">please check your URL or internet connectivity!</p>
   </div>
</div>
</template>
