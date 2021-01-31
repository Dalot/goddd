<template>
  <div class="ui middle aligned center aligned grid">
    <div class="ui placeholder segment">
      <div class="column">
        <div class="ui two column very relaxed stackable grid">
          <div class="column">
            <div class="ui form">
              <div class="field">
                <label>Email</label>
                <div class="ui left icon input">
                  <input type="text" v-model="email" placeholder="Email" />
                  <i class="user icon"></i>
                </div>
                <div class="ui negative message" v-show="emailError.length">
                  <i class="close icon"></i>
                  <div class="body">
                    {{ emailError }}
                  </div>
                </div>
              </div>
              <div class="field">
                <label>Password</label>
                <div class="ui left icon input">
                  <input type="password" v-model="password" />
                  <i class="lock icon"></i>
                </div>
                <div class="ui negative message" v-show="passwordError.length">
                  <i class="close icon"></i>
                  <div class="body">
                    {{ passwordError }}
                  </div>
                </div>
              </div>
              <div class="ui blue submit button" v-on:click="submitLogin">
                Login
              </div>
            </div>
            <div class="ui negative message" v-show="errorShow">
              <i class="close icon"></i>
              <div class="header">
                {{ errorMessage }}
              </div>
            </div>
            <div class="ui success message" v-show="successShow">
              <i class="close icon"></i>
              <div class="header">
                {{ successMessage }}
              </div>
            </div>
          </div>
          <div class="middle aligned column">
            <div class="ui big button">
              <i class="signup icon"></i>
              Sign Up
            </div>
          </div>
        </div>
      </div>
      <div class="ui vertical divider">Or</div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      msg: 'Welcome to Your Vue.js App',
      email: '',
      emailError: '',
      passwordError: '',
      password: '',
      errorMessage: '',
      errorShow: false,
      successMessage: '',
      successShow: false,
    };
  },
  methods: {
    validEmail(email) {
      const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(email);
    },
    checkForm() {
      this.emailError = '';
      this.passwordError = '';

      if (!this.validEmail(this.email)) {
        this.emailError = 'Email is not valid';
      }

      if (!this.email) {
        this.emailError = 'Please insert an email';
      }

      if (this.password.length < 6) {
        this.passwordError = 'Password must be at least 6 chars.';
      }

      if (!this.password) {
        this.passwordError = 'Please insert a password';
      }

      return this.emailError.length === 0 && this.passwordError.length === 0;
    },
    submitLogin() {
      const isValid = this.checkForm();
      if (!isValid) return;

      this.errorShow = false;
      this.successShow = false;
      axios
        .post('http://localhost:8080/login', {
          email: this.email,
          password: this.password,
        })
        .then((response) => {
          const data = response.data;
          console.log(response, data);
          if (data.status === 'error') {
            this.errorMessage = data.message;
            this.errorShow = true;
          } else if (data.status === 'ok') {
            this.successMessage = data.message;
            this.successShow = true;
            localStorage.setItem('token', data.token);
            this.$router.push({ name: '/home' });
          }
        });
    },
  },
};
</script>
<style>
</style>
