<template>
  <div class="page-login">
    <h2 class="ui dividing header">
      Sign up, we'll shatter trello down!<a
        class="anchor"
        id="group-variations"
      ></a>
    </h2>
    <div class="ui centered grid container">
      <div class="nine wide column">
        <div class="ui icon warning message" v-show="signUpFailed">
          <i class="lock icon"></i>
          <div class="content">
            <div class="header">Sign up failed!</div>
            <p>{{ generalErrorMsg }}</p>
          </div>
        </div>
        <div class="ui icon success message" v-show="signUpSucceeded">
          <i class="heart icon"></i>
          <div class="content">
            <div class="header">Sign up succeeded!</div>
          </div>
        </div>
        <div class="ui fluid card">
          <div class="content">
            <form class="ui form" method="POST">
              <div class="field">
                <label>Username</label>
                <input
                  type="text"
                  name="username"
                  v-model="username"
                  placeholder="Username"
                />
              </div>
              <div class="field">
                <label>Email</label>
                <input
                  v-model="email"
                  type="email"
                  name="email"
                  placeholder="johndoe@example.com"
                />
              </div>
              <div class="ui negative message" v-show="errors.email.length">
                <i class="close icon"></i>
                <div class="body">
                  {{ errors.email }}
                </div>
              </div>
              <div class="ui negative message" v-show="errors.username.length">
                <i class="close icon"></i>
                <div class="body">
                  {{ errors.username }}
                </div>
              </div>
              <div class="field">
                <label>Password</label>
                <input
                  type="password"
                  v-model="password"
                  name="password"
                  placeholder="Password"
                />
              </div>
              <div class="ui negative message" v-show="errors.password.length">
                <i class="close icon"></i>
                <div class="body">
                  {{ errors.password }}
                </div>
              </div>
              <div class="field">
                <label>Confirm Password</label>
                <input
                  v-model="confirm_password"
                  type="password"
                  name="confirm_password"
                  placeholder="Password"
                />
              </div>
              <div
                class="ui negative message"
                v-show="errors.confirm_password.length"
              >
                <i class="close icon"></i>
                <div class="body">
                  {{ errors.confirm_password }}
                </div>
              </div>
              <div
                class="ui primary labeled icon button"
                type="submit"
                v-on:click="signup"
              >
                <i class="unlock alternate icon"></i>
                Sign Up
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      email: '',
      username: '',
      password: '',
      confirm_password: '',
      errors: {
        email: '',
        username: '',
        password: '',
        confirm_password: '',
      },
      signUpFailed: false,
      signUpSucceeded: false,
      generalErrorMsg: 'Please check your input!',
    };
  },
  methods: {
    validEmail(email) {
      const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(email);
    },
    checkForm() {
      this.errors.email = '';
      this.errors.password = '';
      this.errors.username = '';
      this.errors.confirm_password = '';

      if (!this.username) {
        this.errors.username = 'Please insert a username';
      }

      if (this.username.length < 3) {
        this.errors.username = 'Must be at least 3 chars';
      }

      if (!this.validEmail(this.email)) {
        this.errors.email = 'Email is not valid';
      }

      if (!this.email) {
        this.errors.email = 'Please insert an email';
      }

      if (this.password.length < 6) {
        this.errors.password = 'Password must be at least 6 chars.';
      }

      if (!this.password) {
        this.errors.password = 'Please insert a password';
      }

      if (this.password !== this.confirm_password) {
        this.errors.confirm_password = 'Passwords do not match';
      }

      // eslint-disable-next-line no-restricted-syntax
      for (const [_, value] of Object.entries(this.errors)) {
        // eslint-disable-next-line no-unused-expressions
        _;
        if (value.length > 0) {
          return false;
        }
      }
      return true;
    },
    signup() {
      this.signUpFailed = false;
      this.signUpSucceeded = false;
      const isValid = this.checkForm();
      if (!isValid) return;

      axios
        .post('http://localhost:8080/signup', {
          email: this.email,
          username: this.username,
          password: this.password,
          confirm_password: this.confirm_password,
        })
        .then((response) => {
          const data = response.data;
          if (data.status === 'error') {
            this.signUpFailed = true;
            this.generalErrorMsg = data.message;
          } else if (data.status === 'ok') {
            this.signUpSucceeded = true;
            this.$router.push('Login');
          }
        });
    },
  },
};
</script>
<style>
</style>
