const input = document.getElementById("input");
function convert(e) {
  if (input.files.length === 0) return;
  const reader = new FileReader();
  reader.addEventListener("load", (event) => {
    if (!event.target) return;
    console.log(event.target.result);
    // console.log(new Uint8Array(event.target.result));
    if (event.target.result instanceof ArrayBuffer) {
      console.log(convertTachiyomi(new Uint8Array(event.target.result)));
    }
    if (typeof event.target.result === "string") {
      const a = convertPaperback(event.target.result).value;
      console.log(a);
      var blob = new Blob([a], { type: "application/gzip" });
      var link = document.createElement("a");
      link.href = window.URL.createObjectURL(blob);
      var fileName = "tachiyomi.proto.gz";
      link.download = fileName;
      link.click();
    }
  });
  // reader.readAsText(e.target.files[0]);
  if (e.target.id === "ty-pb" && input.files[0].type === "application/gzip") {
    reader.readAsArrayBuffer(input.files[0]);
  } else if (
    e.target.id === "pb-ty" &&
    input.files[0].type === "application/json"
  ) {
    reader.readAsText(input.files[0]);
  }
}

const tyPb = document.getElementById("ty-pb");
const pbTy = document.getElementById("pb-ty");
tyPb && tyPb.addEventListener("click", convert);
pbTy && pbTy.addEventListener("click", convert);
