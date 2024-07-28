# visualizer

<!DOCTYPE html>
<html>
<head>
  <style>
    .tab {
      overflow: hidden;
      border: 1px solid #ccc;
      background-color: #f1f1f1;
    }

    .tab button {
      background-color: inherit;
      float: left;
      border: none;
      outline: none;
      cursor: pointer;
      padding: 14px 16px;
      transition: 0.3s;
      font-size: 17px;
    }

    .tab button:hover {
      background-color: #ddd;
    }

    .tab button.active {
      background-color: #ccc;
    }

    .tabcontent {
      display: none;
      padding: 6px 12px;
      border: 1px solid #ccc;
      border-top: none;
    }
  </style>
</head>
<body>

<h2>タブメニュー</h2>

<div class="tab">
  <button class="tablinks" onclick="openTab(event, 'Tab1')">Tab1</button>
  <button class="tablinks" onclick="openTab(event, 'Tab2')">Tab2</button>
</div>

<div id="Tab1" class="tabcontent">
<object data="./LICENSES/MPL-2.0_en.md" type="text/plain" width="100%" height="100%"></object>
</div>

<div id="Tab2" class="tabcontent">
</div>

<script>
function openTab(evt, tabName) {
  var i, tabcontent, tablinks;
  tabcontent = document.getElementsByClassName("tabcontent");
  for (i = 0; i < tabcontent.length; i++) {
    tabcontent[i].style.display = "none";
  }
  tablinks = document.getElementsByClassName("tablinks");
  for (i = 0; i < tablinks.length; i++) {
    tablinks[i].className = tablinks[i].className.replace(" active", "");
  }
  document.getElementById(tabName).style.display = "block";
  evt.currentTarget.className += " active";
}
</script>

</body>
</html>
