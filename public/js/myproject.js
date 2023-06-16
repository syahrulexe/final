let dataProject = [];

function addProject(event){
    event.preventDefault();

    let projectName = document.getElementById("input-project-name").value;
    let startDate = document.getElementById("input-start").value;
    let endDate = document.getElementById("input-end").value;
    let description = document.getElementById("input-description").value;
    let checkboxHtml = document.getElementById("input-check-html").checked ? `<i class="fa-brands fa-html5"></i>` : "";
    let checkboxCss = document.getElementById("input-check-css").checked ? `<i class="fa-brands fa-css3-alt"></i>` : "";
    let checkboxJavascript = document.getElementById("input-check-javascript").checked ? `<i class="fa-brands fa-js"></i>` : "";
    let checkboxJava = document.getElementById("input-check-java").checked ? `<i id="java-logo" class="fa-brands fa-java"></i>` : "";
    let image = document.getElementById("input-project-image").files;

    if (projectName == "") {
        return alert("nama projectnya isi dulu");
      } else if (startDate == "") {
        return alert("start date belom disi");
      } else if (endDate == "") {
        return alert("end date belom disi");
      } else if (description == "") {
        return alert("isi description");
      } else if (image == "") {
        return alert("masukin file gambar");
      } 
    

    image = URL.createObjectURL(image[0]);
    console.log(image);
    
    let project = {
        projectName,
        description,
        checkboxHtml,
        checkboxCss,
        checkboxJavascript,
        checkboxJava,
        distance: getDistanceTime(),
        image,
    };

    
    dataProject.push(project);
    console.log(dataProject);

    renderProject();
}

// get distance time(point of task 5)

function getDistanceTime() {
  let timeStart = new Date(document.getElementById("input-start").value);
  let timeEnd = new Date(document.getElementById("input-end").value);

  let distance =  timeEnd - timeStart;
  console.log(distance);

  let milisecond = 1000 // 1 minute is 1000 milisecond 
  let secondInHours = 3600 // 1 hour is 3600 second
  let hoursInDays = 24 // 1 day is 24 hours
  let daysInWeeks= 7 // 1 weeks is 7 days
  let weeksInMonths = 4 // 1 month is 4 weeks
  let monthsInYears = 12 // 1 year is 12 months

  let distanceYears = Math.floor(  distance / (milisecond * secondInHours * hoursInDays * daysInWeeks * weeksInMonths * monthsInYears));
  let distanceMonths = Math.floor(  distance / (milisecond * secondInHours * hoursInDays * daysInWeeks * weeksInMonths));
  let distanceWeek = Math.floor(  distance / (milisecond * secondInHours * hoursInDays * daysInWeeks));
  let distanceDay = Math.floor(   distance / (milisecond * secondInHours * hoursInDays));

  if (distanceYears >= 2){
    return `Duration : ${distanceYears} Years `;
  } else if (distanceYears == 1){
    return `Duration : ${distanceYears} Year `;
  } else if (distanceMonths >= 2){
    return `Duration : ${distanceMonths} Months`;
  } else if (distanceMonths == 1){
    return `Duration : ${distanceMonths} Month`;
  } else if (distanceWeek >= 2){
    return `Duration ${distanceWeek} Weeks`;
  } else if (distanceWeek == 1){
    return `Duration : ${distanceWeek} Week`;
  } else  if (distanceDay >= 2){
    return `Duration : ${distanceDay} Days`;
  }else  if (distanceDay == 1){
    return `Duration : ${distanceDay} Day`;
  } 
}


function renderProject() {
    document.getElementById("contents").innerHTML = "";

    for (let i = 0; i < dataProject.length; i++) {
        document.getElementById("contents").innerHTML += `
        
                        <div class="card mt-3 shadow p-2 bg-body-tertiary rounded">
                          <img class="img-fluid w-100 object-fit-cover border rounded" style="height: 200px;" src="${dataProject[i].image}" alt="bmth">
                          <h4><a class="text-decoration-none" href="detailproject.html">${dataProject[i].projectName} </a></h4>
                          <p style="font-size: 15px; color: grey">${dataProject[i].distance}</p>
                          <p>${dataProject[i].description}</p>
                          ${dataProject[i].checkboxHtml}
                          ${dataProject[i].checkboxCss}
                          ${dataProject[i].checkboxJavascript}
                          ${dataProject[i].checkboxJava}
                          <div class="mt-3">
                              <button type="button" class="btn btn-dark" style="width: 49%;">edit</button>
                              <button type="button" class="btn btn-dark" style="width: 49%;">delete</button>
                          </div> 
                        </div>  `;
    }
}

function getFullTime(time) {
  
  let monthName = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
  ];

  let date = time.getDate();
  console.log(date);

  let monthIndex = time.getMonth();
  console.log(monthIndex);

  let year = time.getFullYear();
  console.log(year);

  let hours = time.getHours();
  let minutes = time.getMinutes();
  console.log(minutes);

  if (hours <= 9) {
    hours = "0" + hours;
  } else if (minutes <= 9) {
    minutes = "0" + minutes;
  }

  return `Posted at ${date} ${monthName[monthIndex]} ${year} ${hours}:${minutes} WIB `;

}

