<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<script lang="ts">
export default {
    data() {
        return {
            email:"",
            password:"",
        }
    },
    methods:{
        input(i:any){
            // (this as any)[i.target.name] = i.target.value
            console.log(i.target.value)
        },
        submit(){
            console.log(this.email, this.password)
        },
    }
}
</script>
<template>
    <input @input="input" name="email" :value="email"  />
    <input  @input="input" name="password" :value="password"  />
    <button @click="submit" class=" bg-green-500 px-5 py-1 text-white text-md font-semibold rounded-md" >Show Console</button>
</template>
<style lang="scss" scoped>

</style>