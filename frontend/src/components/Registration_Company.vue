<template>
    <div>
        <b-container>
            <b-card bg-variant="light">
                <b-form-group
                        label-cols-lg="2"
                        label="Unternehmen Registration"
                        label-size="lg"
                        label-class="font-weight-bold pt-0"
                        class="mb-0"
                >
                    <b-form-group
                            label-cols-sm="3"
                            label="E-Mail:"
                            label-align-sm="right"
                            label-for="email"
                    >
                        <b-form-input v-model="company.email" type="email" id="email"  placeholder="max@mustermann.de" required ></b-form-input>
                    </b-form-group>

                    <b-form-group
                            label-cols-sm="3"
                            label="Passwort:"
                            label-align-sm="right"
                            label-for="password"
                    >
                        <b-form-input v-model="company.password" type="password" id="password" placeholder="Passwort" required></b-form-input>
                    </b-form-group>

                    <b-form-group
                            label-cols-sm="3"
                            label="Handelsnummer:"
                            label-align-sm="right"
                            label-for="businessNr"
                    >
                        <b-form-input v-model="company.businessNr" id="businessNr" placeholder="0123456789" required></b-form-input>

                    </b-form-group>

                    <b-form-group
                            label-cols-sm="3"
                            label="Name:"
                            label-align-sm="right"
                            label-for="companyName"
                    >
                        <b-form-input v-model="company.companyName" id="companyName" placeholder="Name" required minlength="3" maxlength="30"></b-form-input>
                        <b-form-invalid-feedback id="nameFeedback">
                            Enter at least 3 letters
                        </b-form-invalid-feedback>
                    </b-form-group>

                    <b-form-group
                            label-cols-sm="3"
                            label="Stadt:"
                            label-align-sm="right"
                            label-for="town"
                    >
                        <b-form-input v-model="company.town" id="town" placeholder="Musterstadt" required minlength="3" maxlength="30"></b-form-input>
                    </b-form-group>

                    <b-form-group
                            label-cols-sm="3"
                            label="Straße:"
                            label-align-sm="right"
                            label-for="street"
                    >
                        <b-form-input v-model="company.street" id="street" placeholder="Musterweg 12" required minlength="3" maxlength="30"></b-form-input>
                    </b-form-group>

                    <b-form-group
                            label-cols-sm="3"
                            label="Branche:"
                            label-align-sm="right"
                            label-for="business"
                    >
                        <!-- <b-form-input v-model="company.business" id="business" placeholder="Musterbranche"> </b-form-input> -->
                        <b-form-select v-model="company.selected" id="business" :options="company.options" required></b-form-select>
                    </b-form-group>

                    <b-form-group
                            label-cols-sm="3"
                            label="PayPal Donate Button:"
                            label-align-sm="right"
                            label-for="paypal"
                    >
                        <b-form-input v-model="company.paypal" id="paypal" placeholder="AB2342JJPY5SS" required minlength="3" maxlength="500"></b-form-input>
                        <b-input-group-append style="margin-top: 10px">
                            <b-button variant="outline-danger" v-b-modal.modal-helpPayPal> Button-ID erstellen </b-button>
                        </b-input-group-append>
                    </b-form-group>

                    <b-modal hide-footer id="modal-helpPayPal" size="xl" title="Button erstellen">
                        <HelpPayPal></HelpPayPal>
                    </b-modal>

                    <b-form-group
                            label-cols-sm="3"
                            label="Kurzbeschreibung:"
                            label-align-sm="right"
                            label-for="shortDescription"
                    >
                        <b-form-textarea  v-model="company.shortDescription" id="shortDescription" placeholder="Wir sind..." required
                                      aria-describedby="shortDescriptionInvalidFeedback"  @focus="shortDescriptionFocused = true"
                                      @blur="shortDescriptionFocused = false"
                                      v-focus="shortDescriptionFocused"
                                      :state="shortDescriptionState"></b-form-textarea>
                        <b-form-invalid-feedback v-if="shortDescriptionFocused == true" id="shortDescriptionInvalidFeedback">
                            Gib bitte eine Beschreibung mit mindestens 10 Zeichen an.
                        </b-form-invalid-feedback>
                    </b-form-group>

                    <b-form-group
                            label-cols-sm="3"
                            label="Beschreibung:"
                            label-align-sm="right"
                            label-for="description"
                    >
                        <b-form-textarea v-model="company.description" id="description" placeholder="Wir befinden uns in folgender Situation..." required
                                      aria-describedby="descriptionInvalidFeedback"
                                      @focus="descriptionFocused = true"
                                      @blur="descriptionFocused = false"
                                      v-focus="descriptionFocused"
                                      :state="descriptionState"
                                         rows="5"
                        ></b-form-textarea>
                        <b-form-invalid-feedback  v-if="descriptionFocused == true" id="descriptionInvalidFeedback">
                            Gib bitte eine Beschreibung mit mindestens 50 Zeichen an.
                        </b-form-invalid-feedback>
                    </b-form-group>


                </b-form-group>

                <div class="file-upload-form">
                    Unternehmensfoto hochladen:
                    <b-form-file  no-drop="true" :state="Boolean(company.imageData)" type="file" placeholder="Keine Datei ausgewählt" @change="previewImage" accept="image/*"></b-form-file>
                </div>

                <div class="image-preview" v-if="company.imageData">
                    <b-img class="preview" fluid :src="company.imageData"></b-img>
                </div>

                 <b-button @click="register(company)" style="float:right; margin-top: 1rem">Registrieren</b-button>

            </b-card>

        </b-container>
    </div>
</template>

<script>
    import HelpPayPal from "./HelpPayPal";

    export default {
        name: "Registration_Company",
        components: {HelpPayPal},
        computed: {
            descriptionState() {

                return this.company.description.length > 50 ? true : false
            },
            shortDescriptionState() {
                return this.company.shortDescription.length > 15 ? true : false
            }
        },
        data() {
            return {
                company: {
                    email: '',
                    password: '',
                    businessNr: '',
                    companyName: '',
                    town: '',
                    street: '',
                    business: '',
                    selected: null,
                    options: [
                        { value: 'drinkingPlace', text: 'Trinkstätte' },
                        { value: 'Cosmetics', text: 'Kosmetik' },
                        { value: 'retail', text: 'Einzelhandel' },
                        { value: 'other', text: 'Andere' },
                    ],
                    description: '',
                    shortDescription: '',
                    imageData: "",  // we will store base64 format of image in this string,
                    paypal: "",
                },
                companyJson: JSON.stringify(this.company),
                descriptionFocused: false,
                shortDescriptionFocused: false,
                text: 'default',


            }
        },

        methods: {
            register: function(user){
                console.log("Registering: " + user.companyName + " as " . user.email);
                console.log(JSON.parse(this.company));
            },
            previewImage: function(event) {
                // Reference to the DOM input element
                var input = event.target;
                // Ensure that you have a file before attempting to read it
                if (input.files && input.files[0]) {
                    // create a new FileReader to read this image and convert to base64 format
                    var reader = new FileReader();
                    // Define a callback function to run, when FileReader finishes its job
                    reader.onload = (e) => {
                        // Note: arrow function used here, so that "this.imageData" refers to the imageData of Vue component
                        // Read image as base64 and set to imageData
                        this.company.imageData = e.target.result;
                    }
                    // Start the reader job - read file as a data url (base64 format)
                    reader.readAsDataURL(input.files[0]);
                }
            },
        },

    }
</script>

<style scoped>

</style>