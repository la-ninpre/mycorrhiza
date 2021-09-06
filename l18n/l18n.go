// Code generated by go-localize; DO NOT EDIT.
// This file was generated by robots at
// 2021-09-27 16:30:27.741694276 +0800 +08 m=+0.002576981

package l18n

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

var localizations = map[string]string{
	"en.admin.newuser_create": "Create",
	"en.admin.newuser_title": "New user",
	"en.admin.panel_about": "About this wiki",
	"en.admin.panel_reindex": "Reindex hyphae",
	"en.admin.panel_safe": "Safe things",
	"en.admin.panel_shutdown": "Shutdown wiki",
	"en.admin.panel_title": "Administrative functions",
	"en.admin.panel_unsafe": "Dangerous things",
	"en.admin.panel_updateheader": "Update header links",
	"en.admin.panel_userlist": "User list",
	"en.admin.panel_users": "Manage users",
	"en.admin.user_delete": "Delete",
	"en.admin.user_delete_heading": "Delete user",
	"en.admin.user_delete_tip": "Remove the user from the database. Changes made by the user will be preserved. It will be possible to take this username later.",
	"en.admin.user_delete_warn": "Are you sure you want to delete {{.name}} from the database? This action is irreversible.",
	"en.admin.user_group_heading": "Change group",
	"en.admin.user_title": "User %s",
	"en.admin.user_update": "Update",
	"en.admin.users_actions": "Actions",
	"en.admin.users_create": "Create a new user",
	"en.admin.users_edit": "Edit",
	"en.admin.users_group": "Group",
	"en.admin.users_name": "Name",
	"en.admin.users_notime": "unknown",
	"en.admin.users_password": "Password",
	"en.admin.users_registered": "Registered at",
	"en.admin.users_reindex": "Reindex users",
	"en.admin.users_title": "Manage users",
	"en.auth.cookie_tip": "By submitting this form you give this wiki a permission to store cookies in your browser. It lets the engine associate your edits with you. You will stay logged in until you log out.",
	"en.auth.error_password": "Wrong password.",
	"en.auth.error_telegram": "Could not authorize using Telegram.",
	"en.auth.error_username": "Unknown username.",
	"en.auth.go_back": "Go back",
	"en.auth.go_home": "Go home",
	"en.auth.go_login": "Go to the login page",
	"en.auth.lock_title": "Locked",
	"en.auth.login_button": "Log in",
	"en.auth.login_header": "Log in to {{.name}}",
	"en.auth.login_title": "Login",
	"en.auth.logout_anon": "You cannot log out because you are not logged in.",
	"en.auth.logout_button": "Confirm",
	"en.auth.logout_header": "Log out?",
	"en.auth.logout_title": "Logout?",
	"en.auth.noauth": "Authentication is disabled. You can make edits anonymously.",
	"en.auth.noregister": "Registrations are currently closed. Administrators can make an account for you by hand; contact them.",
	"en.auth.password": "Password",
	"en.auth.password_tip": "The server stores your password in an encrypted form; even administrators cannot read it.",
	"en.auth.register_button": "Register",
	"en.auth.register_header": "Register on {{.name}}",
	"en.auth.register_title": "Register",
	"en.auth.telegram_tip": "You can log in using Telegram. It only works if you have set your @username in Telegram and this username is free on this wiki.",
	"en.auth.try_again": "Try again",
	"en.auth.username": "Username",
	"en.edit.actions": "Actions",
	"en.edit.bold": "Bold",
	"en.edit.bullets": "Bullet list",
	"en.edit.code": "Code block",
	"en.edit.date": "Insert current date",
	"en.edit.heading": "Heading",
	"en.edit.help": "{{.link}} about mycomarkup",
	"en.edit.help_link": "Learn more",
	"en.edit.highlight": "Highlight",
	"en.edit.hr": "Horizontal bar",
	"en.edit.italic": "Italic",
	"en.edit.link": "Link",
	"en.edit.link_title": "Title",
	"en.edit.markup": "Markup",
	"en.edit.mono": "Monospace",
	"en.edit.new_hypha": "You are creating a new hypha.",
	"en.edit.numbers": "Number list",
	"en.edit.preview": "Preview",
	"en.edit.preview_tip": "Note that the hypha hasn't been saved yet. Here's the preview:",
	"en.edit.preview_title": "Preview of %s",
	"en.edit.rocket": "Rocketlink",
	"en.edit.save": "Save",
	"en.edit.selflink": "Link yourself",
	"en.edit.strike": "Strikethrough",
	"en.edit.sub": "Subtext",
	"en.edit.super": "Supertext",
	"en.edit.tag": "Describe your changes:",
	"en.edit.time": "Insert current time",
	"en.edit.title": "Edit %s",
	"en.edit.transclude": "Transclusion",
	"en.edit.underline": "Underline",
	"en.help.attachment": "Attachment",
	"en.help.configuration": "Configuration (for administrators)",
	"en.help.empty_error_line_1": "Try finding a different entry that would help you.",
	"en.help.empty_error_line_2a": "If you want to write this entry by yourself, consider",
	"en.help.empty_error_line_2b": "to Mycorrhiza Wiki directly.",
	"en.help.empty_error_link": "contributing",
	"en.help.empty_error_title": "This entry does not exist!",
	"en.help.entry_not_found": "Entry not found",
	"en.help.hypha": "Hypha",
	"en.help.interface": "Interface",
	"en.help.lock": "Lock",
	"en.help.main": "Main",
	"en.help.mycomarkup": "Mycomarkup",
	"en.help.recent_changes": "Recent changes",
	"en.help.sibling_hyphae": "Sibling hyphae",
	"en.help.special_pages": "Special pages",
	"en.help.telegram": "Telegram authorization",
	"en.help.title": "Help",
	"en.help.top_bar": "Top bar",
	"en.help.topics": "Help topics",
	"en.help.whitelist": "Whitelist",
	"en.ui.about_admins": "Administrators:",
	"en.ui.about_homepage": "Home page:",
	"en.ui.about_hyphae": "See {{.link}} for information about hyphae on this wiki.",
	"en.ui.about_noauth": "This wiki does not use authorization",
	"en.ui.about_title": "About {{.name}}",
	"en.ui.about_usercount": "User count:",
	"en.ui.about_version": "{{.pre}}Mycorrhiza Wiki{{.post}} version:",
	"en.ui.admin_panel": "Admin panel",
	"en.ui.ask_delete": "Delete %s?",
	"en.ui.ask_delete_tip": "In this version of Mycorrhiza Wiki you cannot undelete a deleted hypha but the history can still be accessed.",
	"en.ui.ask_delete_verb": "delete",
	"en.ui.ask_really": "Do you really want to {{.verb}} hypha {{.name}}?",
	"en.ui.ask_rename": "Rename %s",
	"en.ui.ask_unattach": "Unattach %s?",
	"en.ui.ask_unattach_verb": "unattach",
	"en.ui.attach_empty": "This hypha has no attachment, you can upload it here.",
	"en.ui.attach_include": "Include",
	"en.ui.attach_include_tip": "This attachment is an image. To include it in a hypha, use a syntax like this:",
	"en.ui.attach_link": "What are attachments?",
	"en.ui.attach_new": "Attach",
	"en.ui.attach_new_tip": "You can upload a new attachment. Please do not upload too big pictures unless you need to because may not want to wait for big pictures to load.",
	"en.ui.attach_remove": "Unattach",
	"en.ui.attach_remove_button": "Unattach",
	"en.ui.attach_remove_tip": "Please note that you don't have to unattach before uploading a new attachment.",
	"en.ui.attach_size_value": "{{.n}} byte%s",
	"en.ui.attach_size_value+one": "",
	"en.ui.attach_size_value+other": "s",
	"en.ui.attach_stat": "Stat",
	"en.ui.attach_stat_mime": "MIME type:",
	"en.ui.attach_stat_size": "File size:",
	"en.ui.attach_tip": "You can manage the hypha's attachment on this page.",
	"en.ui.attach_title": "Attachment of {{.name}}",
	"en.ui.attach_upload": "Upload",
	"en.ui.attachment_link": "Manage attachment",
	"en.ui.backlinks_desc": "Hyphae which have a link to the selected hypha are listed below.",
	"en.ui.backlinks_link": "{{.n}} backlink%s",
	"en.ui.backlinks_link+one": "",
	"en.ui.backlinks_link+other": "s",
	"en.ui.backlinks_query": "Backlinks to ‘{{.query}}’",
	"en.ui.backlinks_title": "Backlinks to {{.query}}",
	"en.ui.cancel": "Cancel",
	"en.ui.close_dialog": "Close this dialog",
	"en.ui.confirm": "Confirm",
	"en.ui.delete_link": "Delete",
	"en.ui.diff_title": "Diff of {{.name}} at {{.rev}}",
	"en.ui.edit_link": "Edit text",
	"en.ui.error": "Error",
	"en.ui.error_go_back": "Go back to the hypha.",
	"en.ui.error_text_fetch": "Could not fetch text data",
	"en.ui.error_try_again": "Try again",
	"en.ui.header_no_rights": "You must be a moderator to update header links.",
	"en.ui.history_link": "View history",
	"en.ui.history_title": "History of %s",
	"en.ui.list_desc": "This wiki has {{.n}} %s.",
	"en.ui.list_desc+one": "hypha",
	"en.ui.list_desc+other": "hyphae",
	"en.ui.list_heading": "List of hyphae",
	"en.ui.list_title": "List of pages",
	"en.ui.login": "Login",
	"en.ui.media_download": "Download media",
	"en.ui.media_noaudio": "Your browser does not support audio.",
	"en.ui.media_noaudio_link": "Download audio",
	"en.ui.media_novideo": "Your browser does not support video.",
	"en.ui.media_novideo_link": "Download video",
	"en.ui.no_rights": "Not enough rights",
	"en.ui.notexist_heading": "This hypha does not exist",
	"en.ui.notexist_login": "Log in to your account, if you have one",
	"en.ui.notexist_media": "Upload a media",
	"en.ui.notexist_media_tip1": "Upload a picture, a video or an audio. Most common formats can be accessed from the browser, others can be only downloaded afterwards. You can write a description for the media later.",
	"en.ui.notexist_norights": "You are not authorized to create new hyphae. Here is what you can do:",
	"en.ui.notexist_register": "Register a new account",
	"en.ui.notexist_write": "Write a text",
	"en.ui.notexist_write_button": "Create",
	"en.ui.notexist_write_myco": "Mycomarkup",
	"en.ui.notexist_write_tip1": "Write a note, a diary, an article, a story or anything textual using {{.myco}}. Full history of edits to the document will be saved.",
	"en.ui.notexist_write_tip2": "Make sure to follow this wiki's writing conventions if there are any.",
	"en.ui.random_no_hyphae": "There are no hyphae",
	"en.ui.random_no_hyphae_tip": "It is impossible to display a random hypha because the wiki does not contain any hyphae",
	"en.ui.recent_count_post": "recent changes",
	"en.ui.recent_count_pre": "See",
	"en.ui.recent_empty": "Could not find any recent changes.",
	"en.ui.recent_heading": "Recent Changes",
	"en.ui.recent_subscribe": "Subscribe via {{.rss}}, {{.atom}} or {{.json}}",
	"en.ui.recent_subscribe_json": "JSON feed",
	"en.ui.recent_title": "{{.n}} recent change%s",
	"en.ui.recent_title+one": "",
	"en.ui.recent_title+other": "s",
	"en.ui.register": "Register",
	"en.ui.reindex_no_rights": "You must be an admin to reindex hyphae.",
	"en.ui.rename_link": "Rename",
	"en.ui.rename_recurse": "Rename subhyphae too",
	"en.ui.rename_tip": "If you rename this hypha, all incoming links and all relative outcoming links will break. You will also lose all history for the new name. Rename carefully.",
	"en.ui.rename_to": "New name",
	"en.ui.revision_no_text": "This hypha had no text at this revision.",
	"en.ui.revision_title": "{{.name}} at {{.rev}}",
	"en.ui.revision_warning": "Please note that viewing attachments of hyphae is not supported in history for now.",
	"en.ui.search_results_desc": "Every hypha name has been compared with the query. Hyphae that have matched the query are listed below.",
	"en.ui.search_results_query": "Search results for ‘{{.query}}’",
	"en.ui.search_results_title": "Search: {{.query}}",
	"en.ui.sibling_hyphae": "Sibling hyphae",
	"en.ui.subhyphae": "Subhyphae",
	"en.ui.text_link": "View markup",
	"en.ui.title_search": "Search by title",
	"en.ui.users_admins": "Admins",
	"en.ui.users_editors": "Editors",
	"en.ui.users_heading": "List of users",
	"en.ui.users_moderators": "Moderators",
	"en.ui.users_title": "User list",
	"ru.admin.newuser_create": "Создать",
	"ru.admin.newuser_title": "Новый пользователь",
	"ru.admin.panel_about": "Об этой вики",
	"ru.admin.panel_reindex": "Reindex hyphae",
	"ru.admin.panel_safe": "Безопасные действия",
	"ru.admin.panel_shutdown": "Отключить вики",
	"ru.admin.panel_title": "Панель администратора",
	"ru.admin.panel_unsafe": "Опасные действия",
	"ru.admin.panel_updateheader": "Обновить ссылки верхней панели",
	"ru.admin.panel_userlist": "Список пользователей",
	"ru.admin.panel_users": "Менеджер пользователей",
	"ru.admin.user_delete": "Удалить",
	"ru.admin.user_delete_heading": "Удалить пользователя",
	"ru.admin.user_delete_tip": "Удаляет пользователя из базы данных. Правки пользователя будут сохранены. Имя пользователя освободится для повторной регистрации.",
	"ru.admin.user_delete_warn": "Вы уверены, что хотите удалить {{.name}} из базы данных? Это действие нельзя отменить.",
	"ru.admin.user_group_heading": "Изменить группу",
	"ru.admin.user_title": "Пользователь %s",
	"ru.admin.user_update": "Обновить",
	"ru.admin.users_actions": "Действия",
	"ru.admin.users_create": "Создать пользователя",
	"ru.admin.users_edit": "Изменить",
	"ru.admin.users_group": "Группа",
	"ru.admin.users_name": "Имя",
	"ru.admin.users_notime": "неизвестно",
	"ru.admin.users_password": "Пароль",
	"ru.admin.users_registered": "Время создания",
	"ru.admin.users_reindex": "Переиндексировать пользователей",
	"ru.admin.users_title": "Менеджер пользователей",
	"ru.auth.cookie_tip": "Отправляя эту форму, вы разрешаете вики хранить cookie в вашем браузере. Это позволит движку связывать ваши правки с вашей учётной записью. Вы будете авторизованы, пока не выйдете из учётной записи.",
	"ru.auth.error_password": "Неверный пароль.",
	"ru.auth.error_telegram": "Не удалось авторизоваться через Телеграм.",
	"ru.auth.error_username": "Неизвестное имя пользователя.",
	"ru.auth.go_back": "Назад",
	"ru.auth.go_home": "Домой",
	"ru.auth.go_login": "На страницу входа",
	"ru.auth.lock_title": "Доступ закрыт",
	"ru.auth.login_button": "Войти",
	"ru.auth.login_header": "Вход в «{{.name}}»",
	"ru.auth.login_title": "Вход",
	"ru.auth.logout_anon": "Вы не можете выйти, потому что ещё не вошли.",
	"ru.auth.logout_button": "Подтвердить",
	"ru.auth.logout_header": "Выйти?",
	"ru.auth.logout_title": "Выйти?",
	"ru.auth.noauth": "Аутентификация отключена. Вы можете делать правки анонимно.",
	"ru.auth.noregister": "Регистрация в текущее время недоступна. Администраторы могут вручную создать вам учётную запись, свяжитесь с ними.",
	"ru.auth.password": "Пароль",
	"ru.auth.password_tip": "Сервер хранит ваш пароль в зашифрованном виде, даже администраторы не смогут его прочесть.",
	"ru.auth.register_button": "Зарегистрироваться",
	"ru.auth.register_header": "Регистрация на «{{.name}}»",
	"ru.auth.register_title": "Регистрация",
	"ru.auth.telegram_tip": "Вы можете войти с помощью Телеграм. Это сработает, если у вашего профиля есть @имя, и оно не занято в этой вики.",
	"ru.auth.try_again": "Ещё раз",
	"ru.auth.username": "Логин",
	"ru.edit.actions": "Действия",
	"ru.edit.bold": "Жирный",
	"ru.edit.bullets": "Маркир. список",
	"ru.edit.code": "Код-блок",
	"ru.edit.date": "Текущая дата",
	"ru.edit.heading": "Заголовок",
	"ru.edit.help": "{{.link}} о микоразметке",
	"ru.edit.help_link": "Подробнее",
	"ru.edit.highlight": "Выделение",
	"ru.edit.hr": "Гориз. черта",
	"ru.edit.italic": "Курсив",
	"ru.edit.link": "Ссылка",
	"ru.edit.link_title": "Текст",
	"ru.edit.markup": "Разметка",
	"ru.edit.mono": "Моноширинный",
	"ru.edit.new_hypha": "Вы создаёте новую гифу.",
	"ru.edit.numbers": "Нумер. список",
	"ru.edit.preview": "Предпросмотр",
	"ru.edit.preview_tip": "Заметьте, эта гифа ещё не сохранена. Вот её предпросмотр:",
	"ru.edit.preview_title": "Предпросмотр «%s»",
	"ru.edit.rocket": "Ссылка-ракета",
	"ru.edit.save": "Сохранить",
	"ru.edit.selflink": "Ссылка на вас",
	"ru.edit.strike": "Зачёркнутый",
	"ru.edit.sub": "Подстрочный",
	"ru.edit.super": "Надстрочный",
	"ru.edit.tag": "Опишите ваши правки:",
	"ru.edit.time": "Текущее время",
	"ru.edit.title": "Редактирование «%s»",
	"ru.edit.transclude": "Трансклюзия",
	"ru.edit.underline": "Подчеркивание",
	"ru.help.attachment": "Вложение",
	"ru.help.configuration": "Конфигурация (для администраторов)",
	"ru.help.empty_error_line_1": "Попробуйте поискать другую страницу, способную вам помочь.",
	"ru.help.empty_error_line_2a": "Если вы хотите написать эту страницу сами, будем рады вашим правкам в ",
	"ru.help.empty_error_line_2b": "Микоризы.",
	"ru.help.empty_error_link": "репозитории",
	"ru.help.empty_error_title": "Этой страницы не существует!",
	"ru.help.entry_not_found": "Запись не найдена",
	"ru.help.hypha": "Гифа",
	"ru.help.interface": "Интерфейс",
	"ru.help.lock": "Блокировка",
	"ru.help.main": "Введение",
	"ru.help.mycomarkup": "Микоразметка",
	"ru.help.recent_changes": "Недавние изменения",
	"ru.help.sibling_hyphae": "Гифы-сиблинги",
	"ru.help.special_pages": "Специальные страницы",
	"ru.help.telegram": "Авторизация через Телеграм",
	"ru.help.title": "Справка",
	"ru.help.top_bar": "Верхняя панель",
	"ru.help.topics": "Темы справки",
	"ru.help.whitelist": "Белый список",
	"ru.ui.about_admins": "Администраторы:",
	"ru.ui.about_homepage": "Домашняя гифа:",
	"ru.ui.about_hyphae": "См. {{.link}}, чтобы узнать о гифах в этой вики.",
	"ru.ui.about_noauth": "В этой вики нет авторизации",
	"ru.ui.about_title": "О вики «{{.name}}»",
	"ru.ui.about_usercount": "Число пользователей:",
	"ru.ui.about_version": "Версия {{.pre}}Микоризы{{.post}}:",
	"ru.ui.admin_panel": "Администрирование",
	"ru.ui.ask_delete": "Удалить «%s»?",
	"ru.ui.ask_delete_tip": "В этой версии Микоризы нельзя отменить удаление гифы, но её история останется доступной.",
	"ru.ui.ask_delete_verb": "удалить",
	"ru.ui.ask_really": "Вы действительно хотите {{.verb}} гифу «{{.name}}»?",
	"ru.ui.ask_rename": "Переименовать «%s»",
	"ru.ui.ask_unattach": "Открепить «%s»?",
	"ru.ui.ask_unattach_verb": "открепить",
	"ru.ui.attach_empty": "Эта гифа не имеет вложения, здесь вы можете его загрузить.",
	"ru.ui.attach_include": "Добавление",
	"ru.ui.attach_include_tip": "Это вложение – изображение. Чтобы добавить его в текст гифы, используйте синтаксис ниже:",
	"ru.ui.attach_link": "Что такое вложение?",
	"ru.ui.attach_new": "Прикрепить",
	"ru.ui.attach_new_tip": "Вы можете загрузить новое вложение. Пожалуйста, не загружайте слишком большие изображения без необходимости, чтобы впоследствии не ждать её долгую загрузку.",
	"ru.ui.attach_remove": "Открепить",
	"ru.ui.attach_remove_button": "Открепить",
	"ru.ui.attach_remove_tip": "Заметьте, чтобы заменить вложение, вам не нужно его перед этим откреплять.",
	"ru.ui.attach_size_value": "{{.n}} %s",
	"ru.ui.attach_size_value+few": "байта",
	"ru.ui.attach_size_value+many": "байт",
	"ru.ui.attach_size_value+one": "байт",
	"ru.ui.attach_stat": "Свойства",
	"ru.ui.attach_stat_mime": "MIME-тип:",
	"ru.ui.attach_stat_size": "Размер файла:",
	"ru.ui.attach_tip": "На этой странице вы можете управлять вложением.",
	"ru.ui.attach_title": "Вложение «{{.name}}»",
	"ru.ui.attach_upload": "Загрузить",
	"ru.ui.attachment_link": "Вложение",
	"ru.ui.backlinks_desc": "Ниже перечислены гифы, содержащие ссылку на выбранную гифу.",
	"ru.ui.backlinks_link": "{{.n}} %s сюда",
	"ru.ui.backlinks_link+few": "ссылки",
	"ru.ui.backlinks_link+many": "ссылок",
	"ru.ui.backlinks_link+one": "ссылка",
	"ru.ui.backlinks_query": "Обратные ссылки на «{{.query}}»",
	"ru.ui.backlinks_title": "Обратные ссылки на {{.query}}",
	"ru.ui.cancel": "Отмена",
	"ru.ui.close_dialog": "Закрыть этот диалог",
	"ru.ui.confirm": "Применить",
	"ru.ui.delete_link": "Удалить",
	"ru.ui.diff_title": "Разница для «{{.name}}» из {{.rev}}",
	"ru.ui.edit_link": "Редактировать",
	"ru.ui.error": "Ошибка",
	"ru.ui.error_go_back": "Вернуться к гифе.",
	"ru.ui.error_text_fetch": "Не удалось получить текстовые данные",
	"ru.ui.error_try_again": "Попробуйте ещё раз",
	"ru.ui.header_no_rights": "Вы должны быть модератором, чтобы обновить ссылки в заголовке.",
	"ru.ui.history_link": "История",
	"ru.ui.history_title": "История «%s»",
	"ru.ui.list_desc": "В этой вики {{.n}} %s.",
	"ru.ui.list_desc+few": "гифы",
	"ru.ui.list_desc+many": "гиф",
	"ru.ui.list_desc+one": "гифа",
	"ru.ui.list_heading": "Список гиф",
	"ru.ui.list_title": "Список страниц",
	"ru.ui.login": "Войти",
	"ru.ui.media_download": "Скачать медиа",
	"ru.ui.media_noaudio": "Ваш браузер не поддерживает аудио.",
	"ru.ui.media_noaudio_link": "Скачать аудио",
	"ru.ui.media_novideo": "Ваш браузер не поддерживает видео.",
	"ru.ui.media_novideo_link": "Скачать видео",
	"ru.ui.no_rights": "Недостаточно прав",
	"ru.ui.notexist_heading": "Эта гифа не существует",
	"ru.ui.notexist_login": "Войти в свою учётную запись, если она у вас есть",
	"ru.ui.notexist_media": "Загрузить медиа",
	"ru.ui.notexist_media_tip1": "Загрузите изображение, видео или аудио. Распространённые форматы можно просматривать из браузера, остальные – просто скачать. Позже вы можете дописать пояснение к этому медиа.",
	"ru.ui.notexist_norights": "У вас нет прав для создания новых гиф. Вы можете:",
	"ru.ui.notexist_register": "Создать новую учётную запись",
	"ru.ui.notexist_write": "Написать текст",
	"ru.ui.notexist_write_button": "Создать",
	"ru.ui.notexist_write_myco": "микоразметки",
	"ru.ui.notexist_write_tip1": "Напишите заметку, дневник, статью, рассказ или иной текст с помощью {{.myco}}. Сохраняется полная история правок документа.",
	"ru.ui.notexist_write_tip2": "Не забывайте следовать правилам оформления этой вики, если они имеются.",
	"ru.ui.random_no_hyphae": "В этой вики нет гиф",
	"ru.ui.random_no_hyphae_tip": "Невозможно отобразить случайную гифу, потому что вики не содержит ни одной гифы",
	"ru.ui.recent_count_post": "недавних изменений",
	"ru.ui.recent_count_pre": "Отобразить",
	"ru.ui.recent_empty": "Не удалось найти последние изменения.",
	"ru.ui.recent_heading": "Недавние изменения",
	"ru.ui.recent_subscribe": "Подписаться через {{.rss}}, {{.atom}} или {{.json}}",
	"ru.ui.recent_subscribe_json": "JSON-ленту",
	"ru.ui.recent_title": "{{.n}} %s",
	"ru.ui.recent_title+few": "недавних изменения",
	"ru.ui.recent_title+many": "недавних изменений",
	"ru.ui.recent_title+one": "недавнее изменение",
	"ru.ui.register": "Регистрация",
	"ru.ui.reindex_no_rights": "Вы должны быть администратором, чтобы переиндексировать гифы.",
	"ru.ui.rename_link": "Переименовать",
	"ru.ui.rename_recurse": "Также переименовать подгифы",
	"ru.ui.rename_tip": "Если вы переименуете эту гифу, сломаются все ссылки, ведущие на неё, а также исходящие относительные ссылки. Также вы потеряете всю текущую историю для нового названия. Переименовывайте аккуратно.",
	"ru.ui.rename_to": "Новое название",
	"ru.ui.revision_no_text": "В этой ревизии гифы не было текста.",
	"ru.ui.revision_title": "{{.name}} из {{.rev}}",
	"ru.ui.revision_warning": "Обратите внимание, просмотр вложений в истории гифы пока что недоступен.",
	"ru.ui.search_results_desc": "Название каждой из существующих гиф сопоставлено с запросом. Подходящие гифы приведены ниже.",
	"ru.ui.search_results_query": "Результаты поиска для «{{.query}}»",
	"ru.ui.search_results_title": "Поиск: {{.query}}",
	"ru.ui.sibling_hyphae": "Гифы-сиблинги",
	"ru.ui.subhyphae": "Подгифы",
	"ru.ui.text_link": "Посмотреть разметку",
	"ru.ui.title_search": "Поиск по названию",
	"ru.ui.users_admins": "Администраторы",
	"ru.ui.users_editors": "Редакторы",
	"ru.ui.users_heading": "Список пользователей",
	"ru.ui.users_moderators": "Модераторы",
	"ru.ui.users_title": "Список пользователей",
}

type Replacements map[string]interface{}

type Localizer struct {
	Locale	 string
	FallbackLocale string
	Localizations  map[string]string
}

func New(locale string, fallbackLocale string) *Localizer {
	t := &Localizer{Locale: locale, FallbackLocale: fallbackLocale}
	t.Localizations = localizations
	return t
}

func (t Localizer) SetLocales(locale, fallback string) Localizer {
	t.Locale = locale
	t.FallbackLocale = fallback
	return t
}

func (t Localizer) SetLocale(locale string) Localizer {
	t.Locale = locale
	return t
}

func (t Localizer) SetFallbackLocale(fallback string) Localizer {
	t.FallbackLocale = fallback
	return t
}

func (t Localizer) GetWithLocale(locale, key string, replacements ...*Replacements) string {
	str, ok := t.Localizations[t.getLocalizationKey(locale, key)]
	if !ok {
		str, ok = t.Localizations[t.getLocalizationKey(t.FallbackLocale, key)]
		if !ok {
			return key
		}
	}

        // If the str doesn't have any substitutions, no need to
        // template.Execute.
	if strings.Index(str, "}}") == -1 {
                return str
        }

	return t.replace(str, replacements...)
}

func (t Localizer) Get(key string, replacements ...*Replacements) string {
	str := t.GetWithLocale(t.Locale, key, replacements...)
	return str
}

func (t Localizer) getLocalizationKey(locale string, key string) string {
	return fmt.Sprintf("%v.%v", locale, key)
}

func (t Localizer) replace(str string, replacements ...*Replacements) string {
	b := &bytes.Buffer{}
	tmpl, err := template.New("").Parse(str)
	if err != nil {
		return str
	}

	replacementsMerge := Replacements{}
	for _, replacement := range replacements {
		for k, v := range *replacement {
			replacementsMerge[k] = v
		}
	}

	err = template.Must(tmpl, err).Execute(b, replacementsMerge)
	if err != nil {
		return str
	}
	buff := b.String()
	return buff
}
