document.addEventListener("DOMContentLoaded", () => {
  const searchBar = document.getElementById("search-bar");
  const suggestionsList = document.getElementById("suggestions");
  const searchForm = document.getElementById("search-form");

  searchBar.addEventListener(
    "input",
    debounce(async () => {
      const query = searchBar.value.trim();
      if (query.length > 0) {
        const response = await fetch(
          `/search?query=${encodeURIComponent(query)}`,
          {
            headers: { Accept: "application/json" },
          }
        );
        const suggestions = await response.json();
        displaySuggestions(suggestions);
      } else {
        suggestionsList.innerHTML = ""; // Clear suggestions if input is empty
      }
    }, 300)
  );

  function displaySuggestions(suggestions) {
    suggestionsList.innerHTML = "";
    suggestions.forEach(({ name, type }) => {
      const suggestionItem = document.createElement("div");
      suggestionItem.className = "suggestion-item";
      suggestionItem.textContent = `${name} - ${type}`;

      suggestionItem.addEventListener("click", () => {
        searchBar.value = name;
        searchForm.submit();
        suggestionsList.innerHTML = "";
      });
      suggestionsList.appendChild(suggestionItem);
    });
  }
});

function debounce(func, delay) {
  let timeoutId;
  return function (...args) {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => {
      func(...args);
    }, delay);
  };
}

// short cut
document.addEventListener("keydown", (event) => {
  if (event.ctrlKey && event.key === "f") {
    event.preventDefault(); //to prevent default brouse search
    focusSearchBar();
  } else if (event.altKey && event.key === "h") {
    event.preventDefault();
    window.location.href = "/";
  } else if (event.altKey && event.key === "?") {
    openHelp();
  } else if (event.altKey && event.key === "b") {
                event.preventDefault();
                window.history.back();
              }
});

const focusSearchBar = () => {
  const searchBar = document.getElementById("search-bar");
  if (searchBar) {
    searchBar.focus();
  }
};

function openHelp() {
  alert(
    "Shortcut Guide:\n" +
    "Ctrl + F - Focus Search\n" +
    "Alt + H - Home\n" +
    "Alt + ArrowLeft or Alt + B - Back\n" +
    "Alt + shift + ? - Help"
);
}

/*
  
  const debounceSearch = debounce((query) => {
    console.log(`Searching for: ${query}`);
    // Make API call with the search query
  }, 300);
  
  const searchInput = document.getElementById("search-input");
  searchInput.addEventListener("input", (event) => {
    debounceSearch(event.target.value);
  });


  // A function that makes an API call with the search query
function searchHandler(query) {
    // Make an API call with search query
    getSearchResults(query);
}
// A debounce function that takes a function and a delay as parameters
function debounce(func, delay) {
    // A timer variable to track the delay period
    let timer;
    // Return a function that takes arguments
    return function(...args) {
        // Clear the previous timer if any
        clearTimeout(timer);
        // Set a new timer that will execute the function after the delay period
        timer = setTimeout(() => {
            // Apply the function with arguments
            func.apply(this, args);
        }, delay);
    };
}
// A debounced version of the search handler with 500ms delay
const debouncedSearchHandler = debounce(searchHandler, 500);
// Add an event listener to the search bar input
searchBar.addEventListener("input", (event) => {
    // Get the value of the input
    const query = event.target.value;
    // Call the debounced search handler with the query
    debouncedSearchHandler(query);
});*/
