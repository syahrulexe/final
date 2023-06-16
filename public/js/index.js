let hamburgerIsOpen = false;
    
    function openHamburger() {
      let hamburgerNavContainer = document.getElementById(
        "hamburger-nav-container"
      );
    
      if (!hamburgerIsOpen) {
        console.log(hamburgerIsOpen);
        hamburgerNavContainer.style.display = "block";
        hamburgerIsOpen = true;
      } else {
        console.log(hamburgerIsOpen);
        hamburgerNavContainer.style.display = "none";
        hamburgerIsOpen = false;
      }
    }