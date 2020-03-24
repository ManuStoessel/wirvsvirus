<template>
    <div>
        <b-container fluid="xl">
            <b-row>
                <b-col>
            <b-img
            :src="company.images"
            alt="Unternehmensbild"
            fluid>
            </b-img>



                <h2>{{company.companyName}}</h2>
                {{ company.description }}

                <div style="margin-top: 10px">
                    <h2>Kontakt</h2>
                    <p>Adresse: {{company.street}}, {{company.town}}</p>
                    <p>E-Mail: <a :href="'mailto:' + company.email">{{ company.email }} </a></p>
                </div>

            <div align="left">
                <div>
                    <input type="hidden" name="cmd" value="_s-xclick" />
                    <input type="hidden" name="hosted_button_id" :value="company.payPalButtonId" />
                    <input type="image" src="https://www.paypalobjects.com/de_DE/DE/i/btn/btn_donateCC_LG.gif" border="0"
                           name="submit" title="PayPal - Der sichere, einfache Weg online zu bezahlen!" alt="Mit dem PayPal Button spenden"
                           v-b-modal.modal-comment>
                    <img alt="" border="0" src="https://www.paypal.com/de_DE/i/scr/pixel.gif" width="1" height="1" @click="createComment" />
                </div>
            </div>
                </b-col>
            <b-col>
                <h2>Unterstützer</h2>
                <b-card
                        tag="article"
                        class="mb-2"
                >
                    <b-card-text>
                        <div v-for="(comment, index) in company.comments"
                             :key="index">
                            <h4>{{comment.user}}</h4>
                            {{comment.comment}}
                            <hr>
                        </div>

                    </b-card-text>
                </b-card>
            </b-col>
                </b-row>
        </b-container>

        <b-modal hide-footer id="modal-comment"  title="Spende und hinterlasse Deinen Gruß">
            <CommentCreator :company="company"></CommentCreator>
        </b-modal>

    </div>
</template>

<script>
    import CommentCreator from "./comment/CommentCreator";

    export default {
        name: "Registration_Company",

        data() {
            return {
                text: 'default'
            }
        },
        components:{
            CommentCreator
        },
        methods: {
            createComment: function(){

            }
        },
        props: {
            company: Object
        }
    }
</script>

<style scoped>
    b-img {
        width: 90%;
        height: auto;
    }
</style>