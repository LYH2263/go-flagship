async function j(url, opts) {
  const r = await fetch(url, opts);
  const t = await r.text();
  try { return JSON.parse(t); } catch { return { raw: t, status: r.status }; }
}

document.getElementById("upsert-form").addEventListener("submit", async (e) => {
  e.preventDefault();
  const fd = new FormData(e.target);
  const body = {
    key: fd.get("key"),
    enabled: fd.get("enabled") === "on",
    percent: Number(fd.get("percent")),
    description: fd.get("description") || "",
    rules: [{
      attr: fd.get("attr"),
      op: fd.get("op"),
      values: String(fd.get("values") || "").split(",").map(s => s.trim()).filter(Boolean)
    }]
  };
  const out = await j("/api/flags", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body)
  });
  document.getElementById("upsert-out").textContent = JSON.stringify(out, null, 2);
  refresh();
});

document.getElementById("eval-form").addEventListener("submit", async (e) => {
  e.preventDefault();
  const fd = new FormData(e.target);
  let attrs = {};
  try { attrs = JSON.parse(fd.get("attrs")); } catch {}
  const out = await j("/api/eval", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ key: fd.get("key"), attrs })
  });
  document.getElementById("eval-out").textContent = JSON.stringify(out, null, 2);
});

async function refresh() {
  const list = await j("/api/flags");
  const ul = document.getElementById("flags");
  ul.innerHTML = "";
  (Array.isArray(list) ? list : []).forEach((f) => {
    const li = document.createElement("li");
    li.textContent = (f.Key || f.key) + " enabled=" + (f.Enabled ?? f.enabled) + " percent=" + (f.Percent ?? f.percent);
    ul.appendChild(li);
  });
  document.getElementById("stats").textContent = JSON.stringify(await j("/api/stats"), null, 2);
}

document.getElementById("refresh").addEventListener("click", refresh);
refresh();
