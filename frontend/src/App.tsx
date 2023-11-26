import Layout from "./Components/Layout/Layout";
import './App.css'
import SearchBar from './Components/Searchbar/Searchbar';
import CardProd from "./Components/ProdCard/CardProd";
function App() {

  return (


      <Layout>
      <SearchBar />


    <center>

<div className="waskie">
    <p>
    Witaj na naszej porównywarce cen telefonów - miejscu, gdzie znajdziesz najkorzystniejsze oferty na najnowsze modele smartfonów! Jesteśmy tutaj po to, aby ułatwić Ci wybór idealnego telefonu, dostarczając kompleksowe porównania cenowe ze sprawdzonych sklepów internetowych.

Dzięki naszej intuicyjnej platformie porównywawczej, możesz szybko i łatwo znaleźć najlepsze oferty na telefony mobilne, bez konieczności przeszukiwania wielu stron internetowych. Nasza baza danych jest regularnie aktualizowana, zapewniając Ci dostęp do najnowszych informacji o cenach i promocjach.

Znajdziesz u nas pełne specyfikacje techniczne każdego modelu, recenzje użytkowników oraz profesjonalne opinie, co pozwoli Ci dokładnie zorientować się w dostępnych opcjach. Niezależnie od tego, czy szukasz najnowszego flagowca czy też bardziej przystępnego cenowo modelu, nasza porównywarka pomoże Ci znaleźć najlepszą ofertę na rynku.

Dodatkowo, oferujemy informacje o dodatkowych promocjach, rabatach i darmowej dostawie, dzięki czemu możesz maksymalnie zaoszczędzić na zakupie swojego wymarzonego telefonu. Korzystaj z naszej porównywarki cen i ciesz się nowym smartfonem w atrakcyjnej cenie!

Nie trać czasu na przeszukiwanie różnych stron - znajdź najlepszą ofertę już teraz i ułatw sobie zakupy dzięki naszej porównywarce cen telefonów!</p>


</div>



</center>
<div className="products">
<CardProd
  name="telefon"
  price= {123}
imgSrc= "https://image.ceneostatic.pl/data/products/154326309/f-dreame-h12-dual.jpg"
link= "https://www.ceneo.pl/154326309#tag=pp2"
 />
 <CardProd
  name="kawa"
  price= {1233}
imgSrc= "https://image.ceneostatic.pl/data/products/117369540/f-saeco-xelsis-deluxe-sm8780-00-czarny.jpg"
link= "https://www.ceneo.pl/117369540;02514#tag=pp3"
 />
 <CardProd
  name="szczotka"
  price= {13}
imgSrc= "https://image.ceneostatic.pl/data/products/138638988/f-oral-b-io-series-10-white.jpg"
link= "https://www.ceneo.pl/138638988#tag=pp4"
 />
 </div>
      </Layout>

  )
}

export default App
