<template>

    <div class="searchBox">
            <b-row class="searchRow">
                <b-col>
                    Stadt:
                    <input ref="search1" v-model="city">
                </b-col>
                <b-col>
                    Branche:
                    <input ref="search2" v-model="business">
                </b-col>
            </b-row>
            <b-container class="suchergebnisse">
                    <b-card-group deck>
                    <b-row >
                        <b-col sm="4" v-for="(dealer, index) in filteredList" :key="index">

                                <b-card
                                    img-alt="Image"
                                    img-top
                                    tag="article"
                                    style="max-width: 20rem; min-height: 20rem; max-height: 20rem"
                                    class="mb-2 dealer"
                                    @click="openModal(dealer)"
                                    v-b-modal.modal-4
                            >
                                    <b-card-title :title="dealer.companyName"> </b-card-title>
                                    <b-card-img  :src="dealer.images"></b-card-img>
                                <b-card-text>
                                    {{dealer.shortDescription}}
                                </b-card-text>
                            </b-card>
                        </b-col>
                    </b-row>
                </b-card-group>
            </b-container>
            <b-modal hide-footer id="modal-4"  :title="selectedDealer.company">
                <CompanyDetail :company="selectedDealer"></CompanyDetail>
            </b-modal>

    </div>
</template>

<script>
    import CompanyDetail from "../CompanyDetail";
    export default {
        name: "Search",
        props: {
            dealers: Array
        },
        methods:{
            openModal: function(d){
                this.selectedDealer = d;
            },
            search: function(){
                this.filteredList = [];

                if(this.city === "" && this.business === ""){
                    this.filteredList = Array.from(this.dealers);
                }

                else if(this.business !== "" && this.city !== ""){
                    this.searchForBooth();
                }

                else if(this.city !== ""){
                    this.searchForCity();
                }

                else if(this.business !== "" && this.city === ""){
                    this.searchForBusiness();
                }

                else if(this.business === "" && this.city !== ""){
                    this.searchForCity();
                }

            },
            searchForCity: function(){
                for(let i = 0; i < this.dealers.length; i++){
                    if(this.dealers[i].town.includes(this.city)){
                        this.filteredList.push(this.dealers[i]);
                    }
                }
            },
            searchForBusiness: function(){
                console.log(this.business);
                for(let i = 0; i < this.dealers.length; i++){
                    if(this.dealers[i].business.includes(this.business)){
                        this.filteredList.push(this.dealers[i]);
                    }
                }
            },
            searchForBooth: function(){
                for(let i = 0; i < this.dealers.length; i++){
                    if(this.dealers[i].business.includes(this.business) && this.dealers[i].town.includes(this.city)){

                        this.filteredList.push(this.dealers[i]);
                    }
                }
            }
        },
        components:{
            CompanyDetail
        },
        data(){
            return{
                selectedDealer: Object,
                city: "",
                business: "",
                filteredList: []
            }
        },
        watch:{
            city: function(){
                this.search();
            },
            business: function(){
                this.search();
            }
        },
        mounted() {
            this.filteredList = Array.from(this.dealers);
        }
    }
</script>

<style scoped>
    .dealer{
        width: 40rem;
    }

    .dealerShow{
        margin: auto;
    }

    .searchRow{

    }
    .suchergebnisse{
        position: relative;
        top: 1rem
    }
</style>