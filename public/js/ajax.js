const promise = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open("GET", "https://api.npoint.io/6d3869339712f70b9aa4", true);
    //   console.log(xhr);
    xhr.onload = () => {
      if (xhr.status === 200) {
        // We parsing it so it is easier to read in console
        // response vs responseText, the differences are, responseText is an older version, when response is more newer, but the output is still the same/similiar.
        resolve(JSON.parse(xhr.response));
      } else {
        reject("Error loading data.");
      }
    };
    xhr.onerror = () => {
      reject("Network error.");
    };
    xhr.send();
  });
  
  async function getAllTestimonials() {
    const response = await promise;
    //   console.log(response);
  
    let testimonialHTML = "";
    response.forEach(function (item) {
      testimonialHTML += `<div class="col-lg-4 col-md-4 col-sm-6mb-3">
                              <div class="card mt-3 mb-3 shadow mb-5 bg-body-tertiary rounded">
                                  <div class="card-body  p-2 rounded">
                                    <img src="${item.image}" class="img-fluid w-100 object-fit-cover border rounded" style="height: 200px;" />
                                    <p class="text-start mt-2">${item.quote}</p>
                                    <p class="text-end mb-1 fw-bold">- ${item.author}</p>
                                    <p class="text-end mb-1 fw-bold">${item.rating} <i class="fa-solid fa-star"></i></p>
                                    </div>  
                               </div>   
                            </div>
                          `;
    });
  
    document.getElementById("testimonials").innerHTML = testimonialHTML;
  }
  
  getAllTestimonials();
  
  async function getFilteredTestimonials(rating) {
    const response = await promise;
  
    const testimonialFiltered = response.filter((item) => {
      return item.rating === rating;
    });
  
    //   console.log(testimonialFiltered);
  
    let testimonialHTML = "";
  
    if (testimonialFiltered.length === 0) {
      testimonialHTML = "<h1>Data not found!</h1>";
    } else {
      testimonialFiltered.forEach((item) => {
        testimonialHTML += `<div class="col-lg-4 col-md-4 col-sm-6mb-3">
                            <div class="card mt-3 mb-3 shadow mb-5 bg-body-tertiary rounded">
                                <div class="card-body  p-2 rounded">
                                  <img src="${item.image}" class="img-fluid w-100 object-fit-cover border rounded" style="height: 200px;" />
                                  <p class="text-start mt-2">${item.quote}</p>
                                  <p class="text-end mb-1 fw-bold">- ${item.author}</p>
                                  <p class="text-end mb-1 fw-bold">${item.rating} <i class="fa-solid fa-star"></i></p>
                                  </div>  
                            </div>   
                          </div>
                            `;
      });
    }
  
    document.getElementById("testimonials").innerHTML = testimonialHTML;
  }
  