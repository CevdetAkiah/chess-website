// show form when option is selected

const profileOptions = document.querySelector('[role="profileoptions"]');

profileOptions.addEventListener('click', changeOptionFocus);

function changeOptionFocus(e) {
  const targetOption = e.target
  const index = targetOption.tabIndex;
  if (index === 100) {
    index = 0
  }
  // targetForm = targetOption.querySelector(`[role="form"]`);

  console.log(index)
  docForms = document.forms
  
  // want to ignore the logout form
  forms = [docForms[1], docForms[2], docForms[3], docForms[4]]
  form = forms[index - 1]
  form.removeAttribute("hidden")

  // console.log(targetForm)

  // targetForm.removeAttribute("hidden")

}
