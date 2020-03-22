<template>
  <div id="app">
    <b-container>
      <img src="@/assets/Logo_Projekt_01.png" class="headerImage"/>
      <Navbar @changeView="changeView"></Navbar>
      <component v-bind:is="comp" :dealers="users" :company="user"></component>
      <Footer></Footer>
    </b-container>

  </div>
</template>

<script>
import Mainpage from "./components/mainpage/Mainpage";
import Navbar from "./components/mainpage/Navbar";
import Search from "./components/search/Search";
import Profile from "./components/profile/Profile";
import { BIcon, BIconArrowUp, BIconArrowDown } from 'bootstrap-vue'
import Footer from "./components/footer/Footer";


export default {

  name: 'App',
  components: {
      "mainpage": Mainpage,
      "search": Search,
      "profile": Profile,
      Navbar,
      BIcon,
      BIconArrowUp,
      BIconArrowDown,
      Footer
  },
  mounted() {
    document.title = 'Ladenliebe'

  },
  data(){
    return{
        msg: "test",
        comp: "mainpage",
        user: {
          email: "eichenhuette@gmx.org",
          password: "test",
          businessNr: 123456,
          companyName: "Zur Eichenhütte",
          town: "Hannover",
          street: "Edammer Weg 17",
          business: "Bar",
          description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
          shortDescription: "Eine Kleine Bar innerhalb von Irgendwo",
          images: "https://i1.wp.com/sixstarhotels.org/wp-content/uploads/2017/10/Bar-front-No-copyright-e1507122767368.jpg?w=2160&ssl=1",
          payPalButtonId: "1234sadkjas123"
        },
        users:[
          {
            email: "eichenhuette@gmx.org",
            businessNr: 123456,
            companyName: "Zur Eichenhütte",
            town: "Hannover",
            street: "Edammer Weg 17",
            business: "Bar",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Eine Kleine Bar innerhalb von Hannover, welches die Leute zusammenbringen möchte",
            images: "https://i1.wp.com/sixstarhotels.org/wp-content/uploads/2017/10/Bar-front-No-copyright-e1507122767368.jpg?w=2160&ssl=1",
            payPalButtonId: "1234sadkjas123",
            comments: [
                {
                  user: "Philip",
                  comment: "Mein Lieblingslokal"
                },
                {
                  user: "Robin",
                  comment: "Richtig gut"
                },
                {
                  user: "Jannik",
                  comment: "Trinkt sich gut"
                }
            ]
          },
          {
            email: "dieModeArena@web.de",
            businessNr: 123456,
            companyName: "Die Mode Arena",
            town: "Lehrte",
            street: "An der Parkbrücke 2",
            business: "Mode",
            description: "Die Mode Arena ist ein familiengeführtes Modegeschäft für Beratung und Maßschneiderung. ",
            shortDescription: "Ein Familienunternehmen für Mode Beratung und Anpassungen innerhalb von Lehrte",
            images: "https://modegeschaeft-no10.de/wp-content/uploads/2019/04/MG_7949.jpg",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "Heuracles@email",
            businessNr: 123456,
            companyName: "Heuracles Spezialitäten",
            town: "Hamburg",
            street: "Straße 2",
            business: "Lebensmittel",
            description: "Ein Lebensmittelgeschäft für griechische Spezialitäten und Kostproben.",
            shortDescription: "Lebensmittelgeschäft für griechische Spezialitäten und Kostproben.",
            images: "https://cdn.pixabay.com/photo/2016/03/02/20/13/shopping-1232944_960_720.jpg",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "Friseurinkarin@email.com",
            businessNr: 123456,
            companyName: "Friseur Karin",
            town: "Lehrte",
            street: "Im Loor 2",
            business: "Friseur",
            description: "Unserer kleiner Friseurladen fehlt es an Kunden.",
            shortDescription: "kleiner Friseurladen",
            images: "https://cdn.pixabay.com/photo/2018/03/31/04/48/beauty-salon-3277314_960_720.jpg",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "Heuracles@email",
            businessNr: 123456,
            companyName: "An der Ecke",
            town: "Hamburg",
            street: "Straße 2",
            business: "Lebensmittel",
            description: "Ein kleiner Lebensmittelmarkt, mit Blick auf regionale Produkte",
            shortDescription: "Kleine Bar",
            images: "https://cdn.pixabay.com/photo/2015/09/21/14/24/supermarket-949913_960_720.jpg",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "diepinte@gmail.com",
            businessNr: 123456,
            companyName: "Die Pinte",
            town: "Hameln",
            street: "Marktstrasse 34",
            business: "Bar",
            description: "Wir sind eine gemütloiche Tinkstätte im Zentrum von Hameln, durch die momentane Ausganssperre stehen wir ",
            shortDescription: "Kleine Bar innerhalb von Hameln",
            images: "https://cdn.pixabay.com/photo/2016/03/09/09/42/buildings-1245953_960_720.jpg",
            payPalButtonId: "1234sadkjas123",
            comments: []
          },
          {
            email: "goldener adler@email",
            businessNr: 123456,
            companyName: "Goldener Adler",
            town: "Lehrte",
            street: "Straße 2",
            business: "Bar",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar",
            images: "https://cdn.pixabay.com/photo/2016/11/21/13/04/alcoholic-beverages-1845295_960_720.jpg",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "Heuracles@email",
            businessNr: 123456,
            companyName: "Company 2",
            town: "Hamburg",
            street: "Straße 2",
            business: "Lebensmittel",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar",
            images: "https://cdn.pixabay.com/photo/2016/11/29/11/14/business-1869127__340.jpg",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "diehaarprofis@email.de",
            businessNr: 123456,
            companyName: "Die Haarprofis",
            town: "München",
            street: "Straße 2",
            business: "Mode",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar",
            images: "https://cdn.pixabay.com/photo/2016/08/22/19/27/barbershop-1612726__340.jpg",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "Heuracles@email",
            businessNr: 123456,
            companyName: "Company 2",
            town: "Hamburg",
            street: "Straße 2",
            business: "Lebensmittel",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar",
            images: "https://cdn.pixabay.com/photo/2014/09/04/11/03/supermarket-435452__340.jpg",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {email: "email1@email",
            businessNr: 123456,
            companyName: "Die Trinker",
            town: "Hannover",
            street: "Straße 1",
            business: "Bar",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar innerhalb von Irgendwo",
            images: "https://i1.wp.com/sixstarhotels.org/wp-content/uploads/2017/10/Bar-front-No-copyright-e1507122767368.jpg?w=2160&ssl=1",
            payPalButtonId: "1234sadkjas123",
            comments: []
          },
          {
            email: "dieModeArena@email",
            businessNr: 123456,
            companyName: "Company 2",
            town: "Lehrte",
            street: "Straße 2",
            business: "Mode",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar",
            images: "https://modegeschaeft-no10.de/wp-content/uploads/2019/04/MG_7949.jpg",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "Heuracles@email",
            businessNr: 123456,
            companyName: "Company 2",
            town: "Hamburg",
            street: "Straße 2",
            business: "Lebensmittel",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar",
            images: "https://assets.simpleviewcms.com/simpleview/image/fetch/c_fill,f_jpg,h_605,q_65,w_1200/https://media.newmindmedia.com/TellUs/image/%3Ffile%3DKiwi_h_sten_2018_Mari_Wedum_1948746485.jpg%26dh%3D600%26dw%3D800%26t%3D4",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "dieModeArena@email",
            businessNr: 123456,
            companyName: "Company 2",
            town: "Lehrte",
            street: "Straße 2",
            business: "Mode",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar",
            images: "https://i1.wp.com/sixstarhotels.org/wp-content/uploads/2017/10/Bar-front-No-copyright-e1507122767368.jpg?w=2160&ssl=1",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "Heuracles@email",
            businessNr: 123456,
            companyName: "Company 2",
            town: "Hamburg",
            street: "Straße 2",
            business: "Lebensmittel",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar",
            images: "https://www.imago-images.de/bild/st/98055760/m",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "email1@email",
            businessNr: 123456,
            companyName: "Die Trinker",
            town: "Hannover",
            street: "Straße 1",
            business: "Bar",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar innerhalb von Irgendwo",
            images: "https://i1.wp.com/sixstarhotels.org/wp-content/uploads/2017/10/Bar-front-No-copyright-e1507122767368.jpg?w=2160&ssl=1",
            payPalButtonId: "1234sadkjas123",
            comments: []
          },
          {
            email: "dieModeArena@email",
            businessNr: 123456,
            companyName: "Company 2",
            town: "Lehrte",
            street: "Straße 2",
            business: "Mode",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar",
            images: "https://modegeschaeft-no10.de/wp-content/uploads/2019/04/MG_7949.jpg",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "Heuracles@email",
            businessNr: 123456,
            companyName: "Company 2",
            town: "Hamburg",
            street: "Straße 2",
            business: "Lebensmittel",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar",
            images: "https://assets.simpleviewcms.com/simpleview/image/fetch/c_fill,f_jpg,h_605,q_65,w_1200/https://media.newmindmedia.com/TellUs/image/%3Ffile%3DKiwi_h_sten_2018_Mari_Wedum_1948746485.jpg%26dh%3D600%26dw%3D800%26t%3D4",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "dieModeArena@email",
            businessNr: 123456,
            companyName: "Company 2",
            town: "Lehrte",
            street: "Straße 2",
            business: "Mode",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar",
            images: "https://i1.wp.com/sixstarhotels.org/wp-content/uploads/2017/10/Bar-front-No-copyright-e1507122767368.jpg?w=2160&ssl=1",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
          {
            email: "Heuracles@email",
            businessNr: 123456,
            companyName: "Company 2",
            town: "Hamburg",
            street: "Straße 2",
            business: "Lebensmittel",
            description: "Wir sind eine tolle Bar, welche gerne durch die Leute in dieser schwierigen Zeit unterstützt werden möchte",
            shortDescription: "Kleine Bar",
            images: "https://www.imago-images.de/bild/st/98055760/m",
            payPalButtonId: "12sda4sadkjas123",
            comments: []
          },
        ]
    }
  },
  methods:{
      changeView: function(v){
          this.comp = v;
      }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-left: auto;
  margin-right: auto;
  margin-top: 2%;
  width: 100%;
}
.headerImage{
  max-width: 40%;
  height: auto !important;
  float:left;
}
</style>
