package mtproto

const apiLayer = 23

const apiSchema = `
boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;

true#3fedd339 = True;

vector#1cb5c415 {t:Type} # [ t ] = Vector t;

error#c4b9f9bb code:int text:string = Error;

null#56730bcc = Null;

inputPeerEmpty#7f3b18ea = InputPeer;
inputPeerSelf#7da07ec9 = InputPeer;
inputPeerContact#1023dbe8 user_id:int = InputPeer;
inputPeerForeign#9b447325 user_id:int access_hash:long = InputPeer;
inputPeerChat#179be863 chat_id:int = InputPeer;

inputUserEmpty#b98886cf = InputUser;
inputUserSelf#f7c1b13f = InputUser;
inputUserContact#86e94f65 user_id:int = InputUser;
inputUserForeign#655e74ff user_id:int access_hash:long = InputUser;

inputPhoneContact#f392b7f4 client_id:long phone:string first_name:string last_name:string = InputContact;

inputFile#f52ff27f id:long parts:int name:string md5_checksum:string = InputFile;
inputFileBig#fa4f0bb5 id:long parts:int name:string = InputFile;

inputMediaEmpty#9664f57f = InputMedia;
inputMediaUploadedPhoto#2dc53a7d file:InputFile = InputMedia;
inputMediaPhoto#8f2ab2ec id:InputPhoto = InputMedia;
inputMediaGeoPoint#f9c44144 geo_point:InputGeoPoint = InputMedia;
inputMediaContact#a6e45987 phone_number:string first_name:string last_name:string = InputMedia;
inputMediaUploadedVideo#133ad6f6 file:InputFile duration:int w:int h:int mime_type:string = InputMedia;
inputMediaUploadedThumbVideo#9912dabf file:InputFile thumb:InputFile duration:int w:int h:int mime_type:string = InputMedia;
inputMediaVideo#7f023ae6 id:InputVideo = InputMedia;
inputMediaUploadedAudio#4e498cab file:InputFile duration:int mime_type:string = InputMedia;
inputMediaAudio#89938781 id:InputAudio = InputMedia;
inputMediaUploadedDocument#ffe76b78 file:InputFile mime_type:string attributes:Vector<DocumentAttribute> = InputMedia;
inputMediaUploadedThumbDocument#41481486 file:InputFile thumb:InputFile mime_type:string attributes:Vector<DocumentAttribute> = InputMedia;
inputMediaDocument#d184e841 id:InputDocument = InputMedia;

inputChatPhotoEmpty#1ca48f57 = InputChatPhoto;
inputChatUploadedPhoto#94254732 file:InputFile crop:InputPhotoCrop = InputChatPhoto;
inputChatPhoto#b2e1bf08 id:InputPhoto crop:InputPhotoCrop = InputChatPhoto;

inputGeoPointEmpty#e4c123d6 = InputGeoPoint;
inputGeoPoint#f3b7acc9 lat:double long:double = InputGeoPoint;

inputPhotoEmpty#1cd7bf0d = InputPhoto;
inputPhoto#fb95c6c4 id:long access_hash:long = InputPhoto;

inputVideoEmpty#5508ec75 = InputVideo;
inputVideo#ee579652 id:long access_hash:long = InputVideo;

inputFileLocation#14637196 volume_id:long local_id:int secret:long = InputFileLocation;
inputVideoFileLocation#3d0364ec id:long access_hash:long = InputFileLocation;
inputEncryptedFileLocation#f5235d55 id:long access_hash:long = InputFileLocation;
inputAudioFileLocation#74dc404d id:long access_hash:long = InputFileLocation;
inputDocumentFileLocation#4e45abe9 id:long access_hash:long = InputFileLocation;

inputPhotoCropAuto#ade6b004 = InputPhotoCrop;
inputPhotoCrop#d9915325 crop_left:double crop_top:double crop_width:double = InputPhotoCrop;

inputAppEvent#770656a8 time:double type:string peer:long data:string = InputAppEvent;

peerUser#9db1bc6d user_id:int = Peer;
peerChat#bad0e5bb chat_id:int = Peer;

storage.fileUnknown#aa963b05 = storage.FileType;
storage.filePartial#40bc6f52 = storage.FileType;
storage.fileJpeg#7efe0e = storage.FileType;
storage.fileGif#cae1aadf = storage.FileType;
storage.filePng#a4f63c0 = storage.FileType;
storage.filePdf#ae1e508d = storage.FileType;
storage.fileMp3#528a0677 = storage.FileType;
storage.fileMov#4b09ebbc = storage.FileType;
storage.fileMp4#b3cea0e4 = storage.FileType;
storage.fileWebp#1081464c = storage.FileType;

fileLocationUnavailable#7c596b46 volume_id:long local_id:int secret:long = FileLocation;
fileLocation#53d69076 dc_id:int volume_id:long local_id:int secret:long = FileLocation;

userEmpty#200250ba id:int = User;
userSelf#7007b451 id:int first_name:string last_name:string username:string phone:string photo:UserProfilePhoto status:UserStatus inactive:Bool = User;
userContact#cab35e18 id:int first_name:string last_name:string username:string access_hash:long phone:string photo:UserProfilePhoto status:UserStatus = User;
userRequest#d9ccc4ef id:int first_name:string last_name:string username:string access_hash:long phone:string photo:UserProfilePhoto status:UserStatus = User;
userForeign#75cf7a8 id:int first_name:string last_name:string username:string access_hash:long photo:UserProfilePhoto status:UserStatus = User;
userDeleted#d6016d7a id:int first_name:string last_name:string username:string = User;

userProfilePhotoEmpty#4f11bae1 = UserProfilePhoto;
userProfilePhoto#d559d8c8 photo_id:long photo_small:FileLocation photo_big:FileLocation = UserProfilePhoto;

userStatusEmpty#9d05049 = UserStatus;
userStatusOnline#edb93949 expires:int = UserStatus;
userStatusOffline#8c703f was_online:int = UserStatus;
userStatusRecently#e26f42f1 = UserStatus;
userStatusLastWeek#7bf09fc = UserStatus;
userStatusLastMonth#77ebc742 = UserStatus;

chatEmpty#9ba2d800 id:int = Chat;
chat#6e9c9bc7 id:int title:string photo:ChatPhoto participants_count:int date:int left:Bool version:int = Chat;
chatForbidden#fb0ccc41 id:int title:string date:int = Chat;

chatFull#630e61be id:int participants:ChatParticipants chat_photo:Photo notify_settings:PeerNotifySettings = ChatFull;

chatParticipant#c8d7493e user_id:int inviter_id:int date:int = ChatParticipant;

chatParticipantsForbidden#fd2bb8a chat_id:int = ChatParticipants;
chatParticipants#7841b415 chat_id:int admin_id:int participants:Vector<ChatParticipant> version:int = ChatParticipants;

chatPhotoEmpty#37c1011c = ChatPhoto;
chatPhoto#6153276a photo_small:FileLocation photo_big:FileLocation = ChatPhoto;

messageEmpty#83e5de54 id:int = Message;
message#567699b3 flags:int id:int from_id:int to_id:Peer date:int message:string media:MessageMedia = Message;
messageForwarded#a367e716 flags:int id:int fwd_from_id:int fwd_date:int from_id:int to_id:Peer date:int message:string media:MessageMedia = Message;
messageService#1d86f70e flags:int id:int from_id:int to_id:Peer date:int action:MessageAction = Message;

messageMediaEmpty#3ded6320 = MessageMedia;
messageMediaPhoto#c8c45a2a photo:Photo = MessageMedia;
messageMediaVideo#a2d24290 video:Video = MessageMedia;
messageMediaGeo#56e0d474 geo:GeoPoint = MessageMedia;
messageMediaContact#5e7d2f39 phone_number:string first_name:string last_name:string user_id:int = MessageMedia;
messageMediaDocument#2fda2204 document:Document = MessageMedia;
messageMediaAudio#c6b68300 audio:Audio = MessageMedia;

messageActionEmpty#b6aef7b0 = MessageAction;
messageActionChatCreate#a6638b9a title:string users:Vector<int> = MessageAction;
messageActionChatEditTitle#b5a1ce5a title:string = MessageAction;
messageActionChatEditPhoto#7fcb13a8 photo:Photo = MessageAction;
messageActionChatDeletePhoto#95e3fbef = MessageAction;
messageActionChatAddUser#5e3cfc4b user_id:int = MessageAction;
messageActionChatDeleteUser#b2ae9b0c user_id:int = MessageAction;

dialog#ab3a99ac peer:Peer top_message:int unread_count:int notify_settings:PeerNotifySettings = Dialog;

photoEmpty#2331b22d id:long = Photo;
photo#22b56751 id:long access_hash:long user_id:int date:int caption:string geo:GeoPoint sizes:Vector<PhotoSize> = Photo;

photoSizeEmpty#e17e23c type:string = PhotoSize;
photoSize#77bfb61b type:string location:FileLocation w:int h:int size:int = PhotoSize;
photoCachedSize#e9a734fa type:string location:FileLocation w:int h:int bytes:bytes = PhotoSize;

videoEmpty#c10658a8 id:long = Video;
video#388fa391 id:long access_hash:long user_id:int date:int caption:string duration:int mime_type:string size:int thumb:PhotoSize dc_id:int w:int h:int = Video;

geoPointEmpty#1117dd5f = GeoPoint;
geoPoint#2049d70c long:double lat:double = GeoPoint;

auth.checkedPhone#e300cc3b phone_registered:Bool phone_invited:Bool = auth.CheckedPhone;

auth.sentCode#efed51d9 phone_registered:Bool phone_code_hash:string send_call_timeout:int is_password:Bool = auth.SentCode;
auth.sentAppCode#e325edcf phone_registered:Bool phone_code_hash:string send_call_timeout:int is_password:Bool = auth.SentCode;

auth.authorization#f6b673a4 expires:int user:User = auth.Authorization;

auth.exportedAuthorization#df969c2d id:int bytes:bytes = auth.ExportedAuthorization;

inputNotifyPeer#b8bc5b0c peer:InputPeer = InputNotifyPeer;
inputNotifyUsers#193b4417 = InputNotifyPeer;
inputNotifyChats#4a95e84e = InputNotifyPeer;
inputNotifyAll#a429b886 = InputNotifyPeer;

inputPeerNotifyEventsEmpty#f03064d8 = InputPeerNotifyEvents;
inputPeerNotifyEventsAll#e86a2c74 = InputPeerNotifyEvents;

inputPeerNotifySettings#46a2ce98 mute_until:int sound:string show_previews:Bool events_mask:int = InputPeerNotifySettings;

peerNotifyEventsEmpty#add53cb3 = PeerNotifyEvents;
peerNotifyEventsAll#6d1ded88 = PeerNotifyEvents;

peerNotifySettingsEmpty#70a68512 = PeerNotifySettings;
peerNotifySettings#8d5e11ee mute_until:int sound:string show_previews:Bool events_mask:int = PeerNotifySettings;

peerSettings#818426cd flags:# report_spam:flags.0?true = PeerSettings;

wallPaper#ccb03657 id:int title:string sizes:Vector<PhotoSize> color:int = WallPaper;
wallPaperSolid#63117f24 id:int title:string bg_color:int color:int = WallPaper;

inputReportReasonSpam#58dbcab8 = ReportReason;
inputReportReasonViolence#1e22c78d = ReportReason;
inputReportReasonPornography#2e59d922 = ReportReason;
inputReportReasonOther#e1746d0a text:string = ReportReason;

userFull#771095da user:User link:contacts.Link profile_photo:Photo notify_settings:PeerNotifySettings blocked:Bool real_first_name:string real_last_name:string = UserFull;

contact#f911c994 user_id:int mutual:Bool = Contact;

importedContact#d0028438 user_id:int client_id:long = ImportedContact;

contactBlocked#561bc879 user_id:int date:int = ContactBlocked;

contactStatus#d3680c61 user_id:int status:UserStatus = ContactStatus;

contacts.foreignLinkUnknown#133421f8 = contacts.ForeignLink;
contacts.foreignLinkRequested#a7801f47 has_phone:Bool = contacts.ForeignLink;
contacts.foreignLinkMutual#1bea8ce1 = contacts.ForeignLink;

contacts.myLinkEmpty#d22a1c60 = contacts.MyLink;
contacts.myLinkRequested#6c69efee contact:Bool = contacts.MyLink;
contacts.myLinkContact#c240ebd9 = contacts.MyLink;

contacts.link#eccea3f5 my_link:contacts.MyLink foreign_link:contacts.ForeignLink user:User = contacts.Link;

contacts.contactsNotModified#b74ba9d2 = contacts.Contacts;
contacts.contacts#6f8b8cb2 contacts:Vector<Contact> users:Vector<User> = contacts.Contacts;

contacts.importedContacts#ad524315 imported:Vector<ImportedContact> retry_contacts:Vector<long> users:Vector<User> = contacts.ImportedContacts;

contacts.blocked#1c138d15 blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
contacts.blockedSlice#900802a1 count:int blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;

messages.dialogs#15ba6c40 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;
messages.dialogsSlice#71e094f3 count:int dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;

messages.messages#8c718e87 messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
messages.messagesSlice#b446ae3 count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;

messages.statedMessages#969478bb messages:Vector<Message> chats:Vector<Chat> users:Vector<User> pts:int seq:int = messages.StatedMessages;
messages.statedMessagesLinks#3e74f5c6 messages:Vector<Message> chats:Vector<Chat> users:Vector<User> links:Vector<contacts.Link> pts:int seq:int = messages.StatedMessages;

messages.statedMessage#d07ae726 message:Message chats:Vector<Chat> users:Vector<User> pts:int seq:int = messages.StatedMessage;
messages.statedMessageLink#a9af2881 message:Message chats:Vector<Chat> users:Vector<User> links:Vector<contacts.Link> pts:int seq:int = messages.StatedMessage;

messages.sentMessage#d1f4d35c id:int date:int pts:int seq:int = messages.SentMessage;
messages.sentMessageLink#e9db4a3f id:int date:int pts:int seq:int links:Vector<contacts.Link> = messages.SentMessage;

messages.chats#8150cbd8 chats:Vector<Chat> users:Vector<User> = messages.Chats;

messages.chatFull#e5d7d19c full_chat:ChatFull chats:Vector<Chat> users:Vector<User> = messages.ChatFull;

messages.affectedHistory#b7de36f2 pts:int seq:int offset:int = messages.AffectedHistory;

inputMessagesFilterEmpty#57e2f66c = MessagesFilter;
inputMessagesFilterPhotos#9609a51c = MessagesFilter;
inputMessagesFilterVideo#9fc00e65 = MessagesFilter;
inputMessagesFilterPhotoVideo#56e9f0e4 = MessagesFilter;
inputMessagesFilterPhotoVideoDocuments#d95e73bb = MessagesFilter;
inputMessagesFilterDocument#9eddf188 = MessagesFilter;
inputMessagesFilterAudio#cfc87522 = MessagesFilter;
inputMessagesFilterAudioDocuments#5afbf764 = MessagesFilter;
inputMessagesFilterUrl#7ef0dd87 = MessagesFilter;
inputMessagesFilterGif#ffc86587 = MessagesFilter;

updateNewMessage#13abdb3 message:Message pts:int = Update;
updateMessageID#4e90bfd6 id:int random_id:long = Update;
updateReadMessages#c6649e31 messages:Vector<int> pts:int = Update;
updateDeleteMessages#a92bfe26 messages:Vector<int> pts:int = Update;
updateUserTyping#5c486927 user_id:int action:SendMessageAction = Update;
updateChatUserTyping#9a65ea1f chat_id:int user_id:int action:SendMessageAction = Update;
updateChatParticipants#7761198 participants:ChatParticipants = Update;
updateUserStatus#1bfbd823 user_id:int status:UserStatus = Update;
updateUserName#a7332b73 user_id:int first_name:string last_name:string username:string = Update;
updateUserPhoto#95313b0c user_id:int date:int photo:UserProfilePhoto previous:Bool = Update;
updateContactRegistered#2575bbb9 user_id:int date:int = Update;
updateContactLink#51a48a9a user_id:int my_link:contacts.MyLink foreign_link:contacts.ForeignLink = Update;
updateNewAuthorization#8f06529a auth_key_id:long date:int device:string location:string = Update;
updateNewEncryptedMessage#12bcbd9a message:EncryptedMessage qts:int = Update;
updateEncryptedChatTyping#1710f156 chat_id:int = Update;
updateEncryption#b4a2e88d chat:EncryptedChat date:int = Update;
updateEncryptedMessagesRead#38fe25b7 chat_id:int max_date:int date:int = Update;
updateChatParticipantAdd#3a0eeb22 chat_id:int user_id:int inviter_id:int version:int = Update;
updateChatParticipantDelete#6e5f8c22 chat_id:int user_id:int version:int = Update;
updateDcOptions#8e5e9873 dc_options:Vector<DcOption> = Update;
updateUserBlocked#80ece81a user_id:int blocked:Bool = Update;
updateNotifySettings#bec268ef peer:NotifyPeer notify_settings:PeerNotifySettings = Update;
updateServiceNotification#382dd3e4 type:string message:string media:MessageMedia popup:Bool = Update;
updatePrivacy#ee3b272a key:PrivacyKey rules:Vector<PrivacyRule> = Update;
updateUserPhone#12b9417b user_id:int phone:string = Update;

updates.state#a56c2a3e pts:int qts:int date:int seq:int unread_count:int = updates.State;

updates.differenceEmpty#5d75a138 date:int seq:int = updates.Difference;
updates.difference#f49ca0 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> state:updates.State = updates.Difference;
updates.differenceSlice#a8fb1981 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> intermediate_state:updates.State = updates.Difference;

updatesTooLong#e317af7e = Updates;
updateShortMessage#d3f45784 id:int from_id:int message:string pts:int date:int seq:int = Updates;
updateShortChatMessage#2b2fbd4e id:int from_id:int chat_id:int message:string pts:int date:int seq:int = Updates;
updateShort#78d4dec1 update:Update date:int = Updates;
updatesCombined#725b04c3 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq_start:int seq:int = Updates;
updates#74ae4240 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq:int = Updates;

photos.photos#8dca6aa5 photos:Vector<Photo> users:Vector<User> = photos.Photos;
photos.photosSlice#15051f54 count:int photos:Vector<Photo> users:Vector<User> = photos.Photos;

photos.photo#20212ca8 photo:Photo users:Vector<User> = photos.Photo;

upload.file#96a18d5 type:storage.FileType mtime:int bytes:bytes = upload.File;

dcOption#2ec2a43c id:int hostname:string ip_address:string port:int = DcOption;

config#7dae33e0 date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> chat_big_size:int chat_size_max:int broadcast_size_max:int disabled_features:Vector<DisabledFeature> = Config;

nearestDc#8e1a1775 country:string this_dc:int nearest_dc:int = NearestDc;

help.appUpdate#8987f311 id:int critical:Bool url:string text:string = help.AppUpdate;
help.noAppUpdate#c45a6536 = help.AppUpdate;

help.inviteText#18cb9f78 message:string = help.InviteText;

encryptedChatEmpty#ab7ec0a0 id:int = EncryptedChat;
encryptedChatWaiting#3bf703dc id:int access_hash:long date:int admin_id:int participant_id:int = EncryptedChat;
encryptedChatRequested#c878527e id:int access_hash:long date:int admin_id:int participant_id:int g_a:bytes = EncryptedChat;
encryptedChat#fa56ce36 id:int access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long = EncryptedChat;
encryptedChatDiscarded#13d6dd27 id:int = EncryptedChat;

inputEncryptedChat#f141b5e1 chat_id:int access_hash:long = InputEncryptedChat;

encryptedFileEmpty#c21f497e = EncryptedFile;
encryptedFile#4a70994c id:long access_hash:long size:int dc_id:int key_fingerprint:int = EncryptedFile;

inputEncryptedFileEmpty#1837c364 = InputEncryptedFile;
inputEncryptedFileUploaded#64bd0306 id:long parts:int md5_checksum:string key_fingerprint:int = InputEncryptedFile;
inputEncryptedFile#5a17b5e5 id:long access_hash:long = InputEncryptedFile;
inputEncryptedFileBigUploaded#2dc173c8 id:long parts:int key_fingerprint:int = InputEncryptedFile;

encryptedMessage#ed18c118 random_id:long chat_id:int date:int bytes:bytes file:EncryptedFile = EncryptedMessage;
encryptedMessageService#23734b06 random_id:long chat_id:int date:int bytes:bytes = EncryptedMessage;

messages.dhConfigNotModified#c0e24635 random:bytes = messages.DhConfig;
messages.dhConfig#2c221edd g:int p:bytes version:int random:bytes = messages.DhConfig;

messages.sentEncryptedMessage#560f8935 date:int = messages.SentEncryptedMessage;
messages.sentEncryptedFile#9493ff32 date:int file:EncryptedFile = messages.SentEncryptedMessage;

inputAudioEmpty#d95adc84 = InputAudio;
inputAudio#77d440ff id:long access_hash:long = InputAudio;

inputDocumentEmpty#72f0eaae = InputDocument;
inputDocument#18798952 id:long access_hash:long = InputDocument;

audioEmpty#586988d8 id:long = Audio;
audio#c7ac6496 id:long access_hash:long user_id:int date:int duration:int mime_type:string size:int dc_id:int = Audio;

documentEmpty#36f8c871 id:long = Document;
document#f9a39f4f id:long access_hash:long date:int mime_type:string size:int thumb:PhotoSize dc_id:int attributes:Vector<DocumentAttribute> = Document;

help.support#17c6b5f6 phone_number:string user:User = help.Support;

notifyPeer#9fd40bd8 peer:Peer = NotifyPeer;
notifyUsers#b4c83b4c = NotifyPeer;
notifyChats#c007cec3 = NotifyPeer;
notifyAll#74d07c60 = NotifyPeer;

sendMessageTypingAction#16bf744e = SendMessageAction;
sendMessageCancelAction#fd5ec8f5 = SendMessageAction;
sendMessageRecordVideoAction#a187d66f = SendMessageAction;
sendMessageUploadVideoAction#92042ff7 = SendMessageAction;
sendMessageRecordAudioAction#d52f73f7 = SendMessageAction;
sendMessageUploadAudioAction#e6ac8a6f = SendMessageAction;
sendMessageUploadPhotoAction#990a3c1a = SendMessageAction;
sendMessageUploadDocumentAction#8faee98e = SendMessageAction;
sendMessageGeoLocationAction#176f8ba1 = SendMessageAction;
sendMessageChooseContactAction#628cbc6f = SendMessageAction;

contactFound#ea879f95 user_id:int = ContactFound;

contacts.found#566000e results:Vector<ContactFound> users:Vector<User> = contacts.Found;

inputPrivacyKeyStatusTimestamp#4f96cb18 = InputPrivacyKey;

privacyKeyStatusTimestamp#bc2eab30 = PrivacyKey;

inputPrivacyValueAllowContacts#d09e07b = InputPrivacyRule;
inputPrivacyValueAllowAll#184b35ce = InputPrivacyRule;
inputPrivacyValueAllowUsers#131cc67f users:Vector<InputUser> = InputPrivacyRule;
inputPrivacyValueDisallowContacts#ba52007 = InputPrivacyRule;
inputPrivacyValueDisallowAll#d66b66c9 = InputPrivacyRule;
inputPrivacyValueDisallowUsers#90110467 users:Vector<InputUser> = InputPrivacyRule;

privacyValueAllowContacts#fffe1bac = PrivacyRule;
privacyValueAllowAll#65427b82 = PrivacyRule;
privacyValueAllowUsers#4d5bbe0c users:Vector<int> = PrivacyRule;
privacyValueDisallowContacts#f888fa1a = PrivacyRule;
privacyValueDisallowAll#8b73e763 = PrivacyRule;
privacyValueDisallowUsers#c7f49b7 users:Vector<int> = PrivacyRule;

account.privacyRules#554abb6f rules:Vector<PrivacyRule> users:Vector<User> = account.PrivacyRules;

accountDaysTTL#b8d0afdf days:int = AccountDaysTTL;

account.sentChangePhoneCode#a4f58c4c phone_code_hash:string send_call_timeout:int = account.SentChangePhoneCode;

documentAttributeImageSize#6c37c15c w:int h:int = DocumentAttribute;
documentAttributeAnimated#11b58939 = DocumentAttribute;
documentAttributeSticker#fb0a5727 = DocumentAttribute;
documentAttributeVideo#5910cccb duration:int w:int h:int = DocumentAttribute;
documentAttributeAudio#51448e5 duration:int = DocumentAttribute;
documentAttributeFilename#15590068 file_name:string = DocumentAttribute;

messages.stickersNotModified#f1749a22 = messages.Stickers;
messages.stickers#8a8ecd32 hash:string stickers:Vector<Document> = messages.Stickers;

stickerPack#12b299d4 emoticon:string documents:Vector<long> = StickerPack;

messages.allStickersNotModified#e86602c3 = messages.AllStickers;
messages.allStickers#dcef3102 hash:string packs:Vector<StickerPack> documents:Vector<Document> = messages.AllStickers;

disabledFeature#ae636f24 feature:string description:string = DisabledFeature;

---functions---

invokeAfterMsg#cb9f372d {X:Type} msg_id:long query:!X = X;
invokeAfterMsgs#3dc4b4f0 {X:Type} msg_ids:Vector<long> query:!X = X;
initConnection#69796de9 {X:Type} api_id:int device_model:string system_version:string app_version:string lang_code:string query:!X = X;
invokeWithLayer#da9b0d0d {X:Type} layer:int query:!X = X;

auth.checkPhone#6fe51dfb phone_number:string = auth.CheckedPhone;
auth.sendCode#768d5f4d phone_number:string sms_type:int api_id:int api_hash:string lang_code:string = auth.SentCode;
auth.sendCall#3c51564 phone_number:string phone_code_hash:string = Bool;
auth.signUp#1b067634 phone_number:string phone_code_hash:string phone_code:string first_name:string last_name:string = auth.Authorization;
auth.signIn#bcd51581 phone_number:string phone_code_hash:string phone_code:string = auth.Authorization;
auth.logOut#5717da40 = Bool;
auth.resetAuthorizations#9fab0d1a = Bool;
auth.sendInvites#771c1d97 phone_numbers:Vector<string> message:string = Bool;
auth.exportAuthorization#e5bfffcd dc_id:int = auth.ExportedAuthorization;
auth.importAuthorization#e3ef9613 id:int bytes:bytes = auth.Authorization;
auth.bindTempAuthKey#cdd42a05 perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
auth.sendSms#da9f3e8 phone_number:string phone_code_hash:string = Bool;

account.registerDevice#446c712c token_type:int token:string device_model:string system_version:string app_version:string app_sandbox:Bool lang_code:string = Bool;
account.unregisterDevice#65c55b40 token_type:int token:string = Bool;
account.updateNotifySettings#84be5b93 peer:InputNotifyPeer settings:InputPeerNotifySettings = Bool;
account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;
account.resetNotifySettings#db7e1747 = Bool;
account.updateProfile#f0888d68 first_name:string last_name:string = User;
account.updateStatus#6628562c offline:Bool = Bool;
account.getWallPapers#c04cfac2 = Vector<WallPaper>;
account.reportPeer#ae189d5f peer:InputPeer reason:ReportReason = Bool;
account.checkUsername#2714d86c username:string = Bool;
account.updateUsername#3e0bdd7c username:string = User;
account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
account.deleteAccount#418d4e0b reason:string = Bool;
account.getAccountTTL#8fc711d = AccountDaysTTL;
account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
account.sendChangePhoneCode#a407a8f4 phone_number:string = account.SentChangePhoneCode;
account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;
account.updateDeviceLocked#38df3532 period:int = Bool;

users.getUsers#d91a548 id:Vector<InputUser> = Vector<User>;
users.getFullUser#ca30a5b1 id:InputUser = UserFull;

contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
contacts.getContacts#22c6aa08 hash:string = contacts.Contacts;
contacts.importContacts#da30b32d contacts:Vector<InputContact> replace:Bool = contacts.ImportedContacts;
contacts.deleteContact#8e953744 id:InputUser = contacts.Link;
contacts.deleteContacts#59ab389e id:Vector<InputUser> = Bool;
contacts.block#332b49fc id:InputUser = Bool;
contacts.unblock#e54100bd id:InputUser = Bool;
contacts.getBlocked#f57c350f offset:int limit:int = contacts.Blocked;
contacts.exportCard#84e53737 = Vector<int>;
contacts.importCard#4fe196fe export_card:Vector<int> = User;
contacts.search#11f812d8 q:string limit:int = contacts.Found;
contacts.resolveUsername#bf0131c username:string = User;

messages.getMessages#4222fa74 id:Vector<int> = messages.Messages;
messages.getDialogs#eccf1df6 offset:int max_id:int limit:int = messages.Dialogs;
messages.getHistory#92a1df2f peer:InputPeer offset:int max_id:int limit:int = messages.Messages;
messages.search#7e9f2ab peer:InputPeer q:string filter:MessagesFilter min_date:int max_date:int offset:int max_id:int limit:int = messages.Messages;
messages.readHistory#eed884c6 peer:InputPeer max_id:int offset:int read_contents:Bool = messages.AffectedHistory;
messages.deleteHistory#f4f8fb61 peer:InputPeer offset:int = messages.AffectedHistory;
messages.deleteMessages#14f2dd0a id:Vector<int> = Vector<int>;
messages.receivedMessages#28abcb68 max_id:int = Vector<int>;
messages.setTyping#a3825e50 peer:InputPeer action:SendMessageAction = Bool;
messages.sendMessage#4cde0aab peer:InputPeer message:string random_id:long = messages.SentMessage;
messages.sendMedia#a3c85d76 peer:InputPeer media:InputMedia random_id:long = messages.StatedMessage;
messages.forwardMessages#514cd10f peer:InputPeer id:Vector<int> = messages.StatedMessages;
messages.reportSpam#cf1592db peer:InputPeer = Bool;
messages.hideReportSpam#a8f1709b peer:InputPeer = Bool;
messages.getPeerSettings#3672e09c peer:InputPeer = PeerSettings;
messages.getChats#3c6aa187 id:Vector<int> = messages.Chats;
messages.getFullChat#3b831c66 chat_id:int = messages.ChatFull;
messages.editChatTitle#b4bc68b5 chat_id:int title:string = messages.StatedMessage;
messages.editChatPhoto#d881821d chat_id:int photo:InputChatPhoto = messages.StatedMessage;
messages.addChatUser#2ee9ee9e chat_id:int user_id:InputUser fwd_limit:int = messages.StatedMessage;
messages.deleteChatUser#c3c5cd23 chat_id:int user_id:InputUser = messages.StatedMessage;
messages.createChat#419d9aee users:Vector<InputUser> title:string = messages.StatedMessage;
messages.forwardMessage#3f3f4f2 peer:InputPeer id:int random_id:long = messages.StatedMessage;
messages.getDhConfig#26cf8950 version:int random_length:int = messages.DhConfig;
messages.requestEncryption#f64daf43 user_id:InputUser random_id:int g_a:bytes = EncryptedChat;
messages.acceptEncryption#3dbc0415 peer:InputEncryptedChat g_b:bytes key_fingerprint:long = EncryptedChat;
messages.discardEncryption#edd923c5 chat_id:int = Bool;
messages.setEncryptedTyping#791451ed peer:InputEncryptedChat typing:Bool = Bool;
messages.readEncryptedHistory#7f4b690a peer:InputEncryptedChat max_date:int = Bool;
messages.sendEncrypted#a9776773 peer:InputEncryptedChat random_id:long data:bytes = messages.SentEncryptedMessage;
messages.sendEncryptedFile#9a901b66 peer:InputEncryptedChat random_id:long data:bytes file:InputEncryptedFile = messages.SentEncryptedMessage;
messages.sendEncryptedService#32d439a4 peer:InputEncryptedChat random_id:long data:bytes = messages.SentEncryptedMessage;
messages.receivedQueue#55a5bb66 max_qts:int = Vector<long>;
messages.reportEncryptedSpam#4b0c8c0f peer:InputEncryptedChat = Bool;
messages.readMessageContents#354b5bc2 id:Vector<int> = Vector<int>;
messages.getStickers#ae22e045 emoticon:string hash:string = messages.Stickers;
messages.getAllStickers#aa3bc868 hash:string = messages.AllStickers;

updates.getState#edd4882a = updates.State;
updates.getDifference#a041495 pts:int date:int qts:int = updates.Difference;

photos.updateProfilePhoto#eef579a0 id:InputPhoto crop:InputPhotoCrop = UserProfilePhoto;
photos.uploadProfilePhoto#d50f9c88 file:InputFile caption:string geo_point:InputGeoPoint crop:InputPhotoCrop = photos.Photo;
photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
photos.getUserPhotos#b7ee553c user_id:InputUser offset:int max_id:int limit:int = photos.Photos;

upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
upload.getFile#e3a6cfb5 location:InputFileLocation offset:int limit:int = upload.File;
upload.saveBigFilePart#de7b673d file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;

help.getConfig#c4f9186b = Config;
help.getNearestDc#1fb33026 = NearestDc;
help.getAppUpdate#c812ac7e device_model:string system_version:string app_version:string lang_code:string = help.AppUpdate;
help.saveAppLog#6f02f748 events:Vector<InputAppEvent> = Bool;
help.getInviteText#a4a95186 lang_code:string = help.InviteText;
help.getSupport#9cdf08cd = help.Support;
`