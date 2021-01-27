interface Page {
  name: string;
  url: string;
}

function logout({
  event,
  redirectTo,
}: {
  event: Event;
  redirectTo: Page;
}): void {
  const { name, url } = redirectTo;
  localStorage.removeItem('TOKEN');
  localStorage.removeItem('REFRESH_TOKEN');
  window.location.href = url;
}
