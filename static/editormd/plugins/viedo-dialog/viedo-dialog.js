/*!
 * Link dialog plugin for Editor.md
 *
 * @file        viedo-dialog.js
 * @author      pandao
 * @version     1.2.1
 * @updateTime  2015-06-09
 * {@link       https://github.com/pandao/editor.md}
 * @license     MIT
 */

(function() {

    var factory = function (exports) {

		var pluginName   = "video-dialog";

		exports.fn.viedoDialog = function() {

			var _this       = this;
			var cm          = this.cm;
            var editor      = this.editor;
            var settings    = this.settings;
            var selection   = cm.getSelection();
            var lang        = this.lang;
            var linkLang    = lang.dialog.viedo;
            var classPrefix = this.classPrefix;
			var dialogName  = classPrefix + pluginName, dialog;

			cm.focus();

            if (editor.find("." + dialogName).length > 0)
            {
                dialog = editor.find("." + dialogName);
                dialog.find("[data-url]").val("http://");
                dialog.find("[data-title]").val(selection);

                this.dialogShowMask(dialog);
                this.dialogLockScreen();
                dialog.show();
            }
            else
            {
                var dialogHTML = "<div class=\"" + classPrefix + "form\">" + 
                                        "<label>视频地址</label>" + 
                                        "<input type=\"text\" placeholder=\"请输入视频链接\" value=\"\" data-url />" +
                                        // "<br/>" + 
                                        // "<label>" + linkLang.urlTitle + "</label>" + 
                                        "<input type=\"hidden\" value=\"" + selection + "\" data-title />" + 
                                        "<br/>" +
                                    "</div>";

                dialog = this.createDialog({
                    title      : "插入视频框架",
                    width      : 380,
                    height     : 180,
                    content    : dialogHTML,
                    mask       : settings.dialogShowMask,
                    drag       : settings.dialogDraggable,
                    lockScreen : settings.dialogLockScreen,
                    maskStyle  : {
                        opacity         : settings.dialogMaskOpacity,
                        backgroundColor : settings.dialogMaskBgColor
                    },
                    buttons    : {
                        enter  : [lang.buttons.enter, function() {
                            var url   = this.find("[data-url]").val();
                            var title = this.find("[data-title]").val();

                            if (url === "http://" || url === "")
                            {
                                alert(linkLang.urlEmpty);
                                return false;
                            }

                            /*if (title === "")
                            {
                                alert(linkLang.titleEmpty);
                                return false;
                            }*/
                            
                            var str = "[" + title + "](" + url + " \"" + title + "\")";
                            
                            if (title == "")
                            {
                                // str = "<video ddd src=\"" + url + "\" controls\"controls\"></video>";
                                str =`<iframe src="/api/Player?url=${url}" allowfullscreen="true" style="border-radius: 10px;"></iframe>`
                            }                                

                            cm.replaceSelection(str);

                            this.hide().lockScreen(false).hideMask();

                            return false;
                        }],

                        cancel : [lang.buttons.cancel, function() {                                   
                            this.hide().lockScreen(false).hideMask();

                            return false;
                        }]
                    }
                });
			}
		};

	};
    
	// CommonJS/Node.js
	if (typeof require === "function" && typeof exports === "object" && typeof module === "object")
    { 
        module.exports = factory;
    }
	else if (typeof define === "function")  // AMD/CMD/Sea.js
    {
		if (define.amd) { // for Require.js

			define(["editormd"], function(editormd) {
                factory(editormd);
            });

		} else { // for Sea.js
			define(function(require) {
                var editormd = require("../../editormd");
                factory(editormd);
            });
		}
	} 
	else
	{
        factory(window.editormd);
	}

})();
