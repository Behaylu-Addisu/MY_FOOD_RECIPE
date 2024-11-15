<template>

<footer class=" footer "> 
    <div class=" footerTitle m-5">   
         <p class="text-3xl text-center italic tracking-wider"> EHTIOBEAN RECIPES </p>
    </div>
   <div class="container-fluid mx-auto">
        <a class="cursor-pointer" href="">
           <!-- <img src='../assets/images/mylogo.png' alt="logo" class="footerlogo"> -->
        </a>
        <div class="footer-image-container">
            <img src='../assets/images/footerImage1.jpg' alt=" " class="footerImages mx-8 my-3 p-4">
            <img src='../assets/images/footerImage2.jpg' alt=" " class="footerImages mx-8 my-3 p-4">
            <!-- <img src='../assets/images/footerImage3.jpg' alt=" " class="footerImages">
            <img src='../assets/images/footerImage4.jpg' alt=" " class="footerImages"> -->
        </div>
   </div>
   <div class="middle-footer">
        <div class="site-description footer-side -ml-20">
            <p class="description-title text-3xl my-5 "> Discover the Flavors of Ethiopia </p>
            <p class="description-text text-justify px-2 leading-9">
                Explore authentic EHTIOBEAN recipes, from spicy stews to delicious injera, 
                and bring the rich culinary traditions of Ethiopia to your kitchen. 
                Join our community of food lovers and share your own recipes and cooking tips.
            </p>
        </div>
        <div class="page-links footer-side mx-10 allign-middle">
           <ul>
            <li class="bold"> 
                <a href="" class=" py-1 my-1 pl-6 cursor-pointer no-underline hover:underline"> Home </a>
            </li>
            <li> 
                <a href="" class=" py-1 my-1 pl-6 cursor-pointer no-underline hover:underline"> About Us</a>
            </li>
           </ul>
        </div>
        <div class="socialmedia-links footer-side">
            <p class="contactus text-3xl"> CONTACT US </p>
            <ul>
                <li class=" my-1 py-1 pl-4 hover:underline cursor-pointer"> instagram </li>
                <li class=" my-1 py-1 pl-4 hover:underline cursor-pointer"> facebook </li>
                <li class=" my-1 py-1 pl-4 hover:underline cursor-pointer"> telegram </li>
                <li class=" my-1 py-1 pl-4 hover:underline cursor-pointer"> twiter </li>
            </ul>
        </div>
   </div>
   <p class="copyright text-center mt-4 traking-wider"> © 2024 EHTIOBEAN Recipes. All rights reserved.</p>
</footer>

</template>
<script>
    export default {

    }
</script>
<style scoped>
.footer{
    position: relative;
    bottom: 0;
    left: 0;
    right: 0;
    width: 100%;
    background-image: url("../assets/images/footerBackground1.jpg");
    background-repeat: no-repeat;
    background-position: center;
    background-size: cover;
    min-height: 100px;
    padding: 20px 50px;
    opacity: 0.9;
    color: #fff;

}
.footerTitle{
    margin: 20px auto;
}
.container-fluid{
  display: flex;
  justify-content: space-between;
}
.footerlogo{
    width: 80px;
    height: 80px;
    border-radius: 40px;
    position: absolute;
    left: 0;
    top: 0;
    margin-right: 40px;
}
.footerImages{
    width: 250px;
    height: 200px;
    display: inline-block;
}
.footer-image-container {
    justify-content: center;
    align-items: center;
    margin: auto;
}
.middle-footer{
    display: flex;
    justify-content: space-evenly;
    margin: 10px auto;
}
.footer-side{
    max-width: 50%;
}
</style>