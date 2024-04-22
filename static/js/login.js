document.addEventListener("DOMContentLoaded", function () {
  const firebaseConfig = {
    apiKey: "AIzaSyBOksfyE0unI87E533LWTKV6jpW-F1dqDA",
    authDomain: "i9auth.firebaseapp.com",
    projectId: "i9auth",
    storageBucket: "i9auth.appspot.com",
    messagingSenderId: "284108905250",
    appId: "1:284108905250:web:b866b21949e149af7ca37b",
  };

  firebase.initializeApp(firebaseConfig);
  console.log(firebase.auth().signInWithEmailAndPassword);

  window.login = function () {
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    firebase
      .auth()
      .signInWithEmailAndPassword(email, password)
      .then((userCredential) => {
        console.log(userCredential)
      })
    //   .then((idToken) => {
    //     return postIdTokenToVerifyToken("/verifyToken", idToken);
    //   })
    //   .then(() => {
    //     window.location.assign("/");
    //   })
    //   .catch((error) => {
    //     console.error("Authentication failed:", error);
    //   });

    function postIdTokenToVerifyToken(url, idToken) {
      return fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ idToken }),
      });
    }
  };
});
