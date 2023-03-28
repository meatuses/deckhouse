import FormatError from './FormatError.js';

const RELOAD_CHAIN_THRESHOLD_NUM = 3;
const RELOAD_CHAIN_THRESHOLD_SECONDS = RELOAD_CHAIN_THRESHOLD_NUM * 60;

/*
type FlashEventAdd = {
  key: string,
  text: string,
  type: string
  key?: string,
  isReloader?: Boolean
}

type FlashEventClose = {
  key: string,
}
*/

// STATELESS SERVICE
export default class FlashMessagesService {
  static show(type, text, timeout, key) {
    this.$flashMessages.emit('Flash::add', { type: type, text: text, timeout: timeout, key: key });
  }

  static close(key) {
    this.$flashMessages.emit('Flash::close', key);
  }

  static reload(text) {
    var self = this;
    setTimeout(() => {
      self.show('warning', `ALERT\n ${text}\n WILL RELOAD IN 3 SECONDS!!!`, 0, 'forced_reload_a');
    }, 1);
    setTimeout(() => {
      self.show('warning', 'Reloading...', 0, 'forced_reload_b');
    }, 2500);
    setTimeout(() => {
      location.reload();
    }, 3000);
  }

  static isInReloadChain() {
    try {
      var recentReloadsStr = localStorage.getItem('recentReloads');
      var recentReloads = recentReloadsStr ? JSON.parse(recentReloadsStr) : [];
      var now = Math.round(new Date() / 1000);
      recentReloads.push(now)
      recentReloads = recentReloads.filter((t) => { return t >= now - RELOAD_CHAIN_THRESHOLD_SECONDS; })
      localStorage.setItem('recentReloads', JSON.stringify(recentReloads));
      return recentReloads.length >= RELOAD_CHAIN_THRESHOLD_NUM;

    } catch (error) {
      if (!(error instanceof SyntaxError)) throw error;
      console.log(error);
      localStorage.setItem('recentReloads', JSON.stringify('[]'));
      return false;
    }
  }

  static proposeReload(text, isNotThreateningReloadChain) {
    if (this.initiatedReload || this.initiatedReloadChainEscape) return;

    if (!isNotThreateningReloadChain && this.isInReloadChain()) {
      // Reloads chain escape. Currently, simpliest: just don't automatically reload, user can always press F5.
      this.initiatedReloadChainEscape = true;
      this.$flashMessages.emit('Flash::add', {
          type: 'error',
          text: `Reload chain stopped after ${RELOAD_CHAIN_THRESHOLD_NUM} reloads in less than ${RELOAD_CHAIN_THRESHOLD_SECONDS} seconds. [displayed on ${moment(new Date()).format('HH:mm:ss')}]`,
          key: 'reload',
          isReloader: false
        }
      );

    } else {
      this.initiatedReload = true;
      return this.reload(text);
    }
  }

  static wrapAction(subject, actionName, flashMessage, flashKey, flashErrorType) {
    var self = this;
    var oldAction = subject[actionName];
    subject[actionName] = function() {
      var actionFlashKey = flashKey || this.klassName;
      return oldAction.apply(this, arguments).then((resp) => {
        self.close(actionFlashKey);
        return resp;
      }).catch((error) => {
        self.show(flashErrorType || 'error', flashMessage + ": " + FormatError(error), 0, actionFlashKey);
        return Promise.reject(error);
      });
    };
  }
}