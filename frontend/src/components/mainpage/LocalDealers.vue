<template>
    <div>
        <h1>Deine lokalen Geschäfte</h1>
        <b-row class="dealerShow">
            <b-col v-for="(dealer, index) in dealers.slice(0,5)" :key="index">
                <b-card
                        :title="dealer.companyName"
                        :img-src="dealer.images"
                        img-alt="Image"
                        img-top
                        tag="article"
                        class="mb-2"
                        @click="openModal(dealer)"
                        v-b-modal.modal-3
                >
                    <b-card-text>
                        {{dealer.shortDescription}}
                    </b-card-text>
                </b-card>
            </b-col>
        </b-row>
        <b-modal hide-footer id="modal-3"  :title="selectedDealer.company">
            <CompanyDetail :company="selectedDealer"></CompanyDetail>
        </b-modal>
    </div>
</template>

<script>
    import CompanyDetail from "../CompanyDetail";

    export default {

        name: "LocalDealers",
        components: {
            CompanyDetail
        },
        methods: {
            getLocalDealers: function(){

            },
            openModal: function(d){
                this.selectedDealer = d;
                console.log(d);
            }
        },
        created() {
            this.getLocalDealers();
        },
        props: {
            dealers : Array
        },
        data(){
            return{
                selectedDealer: Object
            }
        }
    }
</script>

<style scoped>
    .dealerShow{
        margin-left: 2%;
    }
    b-card {
        position: relative;
        left: 2rem;
        right: 2rem;
        top: 2rem;
        bottom: 2rem;
        max-width: 20rem;
        min-width: 20rem;
        min-height: 20rem;
        max-height: 20rem
    }
</style>