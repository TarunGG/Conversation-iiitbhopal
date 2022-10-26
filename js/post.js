 console.log('working')
      let body=document.querySelector('body')
      let btn=document.getElementById('goBottomArrow')
      window.onscroll=()=>{
        if(window.scrollY+window.screen.height<body.clientHeight-50){
          btn.style.visibility="visible"
        }
        else{
          btn.style.visibility="hidden"
        }
      }
