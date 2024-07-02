const input = document.getElementById("input");
function convert(e) {
  if (input.files.length === 0) return;
  const reader = new FileReader();
  reader.addEventListener("load", (event) => {
    if (!event.target) return;
    console.log(event.target.result);
    // console.log(new Uint8Array(event.target.result));
    let value, type, fileName;
    if (event.target.result instanceof ArrayBuffer) {
      const converted = convertTachiyomi(new Uint8Array(event.target.result));
      console.log("tachiyomi converted", converted);
      value = [converted.value];
      type = "application/json";
      fileName = "paperback.json";
    } else if (typeof event.target.result === "string") {
      const converted = convertPaperback(event.target.result);
      console.log(converted);
      value = [converted.value];
      type = "application/gzip";
      fileName = "tachiyomi.proto.gz";
    }

    var blob = new Blob(value, { type });
    var link = document.createElement("a");
    link.href = window.URL.createObjectURL(blob);
    link.download = fileName;
    link.click();
  });
  // reader.readAsText(e.target.files[0]);
  if (
    e.target.id === "ty-pb" &&
    ["application/gzip", "application/x-gzip"].includes(input.files[0].type)
  ) {
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
