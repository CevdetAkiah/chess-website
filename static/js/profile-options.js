// show form when option is selected

const profileOptions = document.querySelector('[role="profileoptions"]');

profileOptions.addEventListener('click', changeOptionFocus);

function changeOptionFocus(e) {
  const targetOption = e.target
  const index = targetOption.tabIndex;
  // targetForm = targetOption.querySelector(`[role="form"]`);

  console.log(index)
  forms = document.forms
  form = forms[index]
  form.removeAttribute("hidden")

  // console.log(targetForm)

  // targetForm.removeAttribute("hidden")

}
