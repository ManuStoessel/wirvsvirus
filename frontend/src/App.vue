<template>
  <div id="app">
    <Navbar @changeView="changeView"></Navbar>
    <component v-bind:is="comp" :dealers="users" :company="user"></component>
  </div>
</template>

<script>
import Mainpage from "./components/mainpage/Mainpage";
import Navbar from "./components/mainpage/Navbar";
import Search from "./components/search/Search";
import Profile from "./components/profile/Profile";

export default {
  name: 'App',
  components: {
      "mainpage": Mainpage,
      "search": Search,
      "profile": Profile,
      Navbar,
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
          payPalButtonId: "1234sadkjas123",
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
            shortDescription: "Eine Kleine Bar innerhalb von Irgendwo",
            images: "https://i1.wp.com/sixstarhotels.org/wp-content/uploads/2017/10/Bar-front-No-copyright-e1507122767368.jpg?w=2160&ssl=1",
            payPalButtonId: "1234sadkjas123",
          },
          {
            email: "dieModeArena@web.de",
            businessNr: 123456,
            companyName: "Die Mode Arena",
            town: "Lehrte",
            street: "An der Parkbrücke 2",
            business: "Mode",
            description: "Die Mode Arena ist ein familiengeführtes Modegeschäft für Beratung und Maßschneiderung. ",
            shortDescription: "Ein Familienunternehmen für Mode Beratung und Anpassungen",
            images: "https://modegeschaeft-no10.de/wp-content/uploads/2019/04/MG_7949.jpg",
            payPalButtonId: "12sda4sadkjas123",
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
            images: "https://assets.simpleviewcms.com/simpleview/image/fetch/c_fill,f_jpg,h_605,q_65,w_1200/https://media.newmindmedia.com/TellUs/image/%3Ffile%3DKiwi_h_sten_2018_Mari_Wedum_1948746485.jpg%26dh%3D600%26dw%3D800%26t%3D4",
            payPalButtonId: "12sda4sadkjas123",
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
            images: "https://i1.wp.com/sixstarhotels.org/wp-content/uploads/2017/10/Bar-front-No-copyright-e1507122767368.jpg?w=2160&ssl=1",
            payPalButtonId: "1234sadkjas123",
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
  margin-top: 5%;
  width: 100%;

}
</style>
