<template>
    <div>

        <b-form-group

                label=""
                label-size="lg"
                label-class="font-weight-bold pt-0"
                class="mb-0"
        >
            <b-form-group
                    label-cols-sm="2"
                    label="Spende:"
                    label-align-sm="right"
                    label-for="donation"
            >
                <b-form-input v-model="comment.donation" type="text" id="donation"  placeholder="Deine Spende über PayPal (€)" required ></b-form-input>
            </b-form-group>

            <b-form-group
                    label-cols-sm="2"
                    label="Name:"
                    label-align-sm="right"
                    label-for="username"
            >
                <b-form-input v-model="comment.user" type="text" id="username"  placeholder="Dein Name" required ></b-form-input>
            </b-form-group>

            <b-form-group
                    label-cols-sm="2"
                    label="Text:"
                    label-align-sm="right"
                    label-for="description"
            >
                <b-form-textarea v-model="comment.comment" id="description" placeholder="Haltet durch, wir sind mit euch ..."
                                 aria-describedby="descriptionInvalidFeedback"
                                 @focus="descriptionFocused = true"
                                 @blur="descriptionFocused = false"
                                 v-focus="descriptionFocused"
                                 rows="5"
                ></b-form-textarea>
                <b-form-invalid-feedback  id="descriptionInvalidFeedback">
                    Gib bitte eine Beschreibung mit mindestens 50 Zeichen an.
                </b-form-invalid-feedback>
            </b-form-group>

        </b-form-group>
        <b-btn @click="createComment" style="float:right;">Hochladen</b-btn>
    </div>
</template>

<script>
    export default {
        name: "CommentCreator",
        props:{
            company: Object
        },
        data(){
            return {
                comment: {
                    user: "",
                    comment: ""
                }
            }
        },
        methods:{
            createComment: function(){
                console.log(this.company)
                this.company.comments.push(this.comment);
                this.close();
            },
            close: function () {
                this.$bvModal.hide("modal-comment");
            }
        }
    }
</script>

<style scoped>

</style>