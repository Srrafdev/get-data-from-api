const searchValue = document.getElementById("inpSrsh");
const suggs = document.getElementById("suggs");

searchValue.addEventListener("input", async () => {
   
  let response = await fetch('/api/search', {
          method: 'POST',
          headers: {
              "Content-Type": "application/json",
          },
          body: JSON.stringify({target:searchValue.value})
      })
      let data = await response.json()
      suggs.style.display = "block"
      
      if (data == null || searchValue.value === "" ){
         suggs.style.display = "none"
        return
      }
      console.log(data);
      
      let ineer = "";
      console.log(data);
      
      if(data.name){
        for(let val of data.name){
          let idVal = val.split("++")
          ineer += `<p>Artist : <a href="http://localhost:8080/GetMore?submit=${idVal[0]}">${idVal[1]}</a></p>`
        }
      }
      if(data.locations){
        for(let val of data.locations){
           let idVal = val.split("++")
          ineer += `<p>Location : <a href="http://localhost:8080/GetMore?submit=${idVal[0]}">${idVal[2]}: ${idVal[1]}</a></p>`
        }
      }
      if(data.member){
        for(let val of data.member){
           let idVal = val.split("++")
          ineer += `<p>Member : <a href="http://localhost:8080/GetMore?submit=${idVal[0]}">${idVal[1]}</a></p>`
        }
      }
      if(data.first_album){
        for(let val of data.first_album){
           let idVal = val.split("++")
          ineer += `<p>First Album : <a href="http://localhost:8080/GetMore?submit=${idVal[0]}">${idVal[2]}: ${idVal[1]}</a></p>`
        }
      }
      if(data.creation_date){
        for(let val of data.creation_date){
           let idVal = val.split("++")
          ineer += `<p>Creation Date : <a href="http://localhost:8080/GetMore?submit=${idVal[0]}">${idVal[2]}: ${idVal[1]}</a></p>`
        }
      }
      suggs.innerHTML = ineer

      if((data.name == null && data.member == null && data.locations == null && data.first_album == null && data.creation_date == null)){
         suggs.innerHTML = "<p>Not found</p>"
      }
})

function toggleSidebar(){
  const sidebar = document.getElementById('sidebar');
  if (sidebar) {
      sidebar.classList.toggle('show');
  } else {
      console.error('Sidebar element not found');
  }
}