<template>
    <div>

        <h1>Deine lokalen Geschäfte</h1>

        <b-row class="dealerShow">

            <b-col v-for="(dealer, index) in dealers.slice(0,6)" :key="index">

                <b-container>
                <b-card
                        img-alt="Image"
                        img-top
                        tag="article"
                        style="max-width: 20rem; min-height: 20rem; max-height: 20rem"
                        class="mb-2 localDealer"
                        @click="openModal(dealer)"
                        v-b-modal.modal-3
                >
                    <b-card-title :title="dealer.companyName"> </b-card-title>
                    <b-card-img  :src="dealer.images"></b-card-img>
                    <b-card-text>
                        {{dealer.shortDescription}}
                    </b-card-text>
                </b-card>
            </b-container>


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
                console.log(this.dealers.comments)
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
        margin-bottom: 5%;
    }
    .localDealer {
        max-width: 20rem;
        min-width: 20rem;
        min-height: 20rem;
        max-height: 20rem
    }
</style>